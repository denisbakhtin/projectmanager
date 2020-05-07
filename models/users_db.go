package models

import (
	"fmt"
	"strings"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/jinzhu/gorm"
)

//UsersDB is a user db repository
var UsersDB UsersRepository

func init() {
	UsersDB = usersRepository{}
}

//UsersRepository is a users repository interface
type UsersRepository interface {
	GetAll() ([]User, error)
	Get(id interface{}) (User, error)
	GetByEmail(email string) (User, error)
	UpdateStatus(user User) (User, error)
	UpdateAccount(vm AccountVM, user User) (User, error)
	Login(vm LoginVM) (User, error)
	Activate(vm ActivateVM) (User, error)
	Register(vm RegisterVM) (User, error)
	Forgot(vm ForgotVM) (User, error)
	ResetPassword(vm ResetVM) (User, error)
	Summary() (UsersSummaryVM, error)
}

//usersRepository is a repository of users
type usersRepository struct{}

//GetAll returns all users owned by specified user
func (tr usersRepository) GetAll() ([]User, error) {
	var users []User
	err := db.Order("id asc").Find(&users).Error
	return users, err
}

//Get fetches a user by its id
func (tr usersRepository) Get(id interface{}) (User, error) {
	user := User{}
	err := db.First(&user, id).Error
	return user, err
}

//GetByEmail fetches a user by email
func (tr usersRepository) GetByEmail(email string) (User, error) {
	user := User{}
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}

//UpdateStatus updates user status in db
func (tr usersRepository) UpdateStatus(user User) (User, error) {
	err := db.Model(&user).UpdateColumn("status", user.Status).Error
	return user, err
}

//UpdateAccount updates user name & password in db
func (tr usersRepository) UpdateAccount(vm AccountVM, user User) (User, error) {
	user.Name = vm.Name
	if len(vm.CurrentPassword) > 0 && len(vm.NewPassword) > 0 {
		if !user.HasPassword(vm.CurrentPassword) {
			return User{}, fmt.Errorf("Wrong current password")
		}
		user.PasswordHash = helpers.CreatePasswordHash(vm.NewPassword)
	}
	if err := db.Save(&user).Error; err != nil {
		return User{}, err
	}
	if err := user.CreateJWTToken(); err != nil {
		return User{}, err
	}

	return user, nil
}

//Login checks if credentials are valid for logging in
func (tr usersRepository) Login(vm LoginVM) (User, error) {
	user := User{}
	err := db.Where("email = ?", helpers.NormalizeEmail(vm.Email)).First(&user).Error
	switch {
	case err != nil && gorm.IsRecordNotFoundError(err):
		return User{}, fmt.Errorf("Wrong email or password")
	case err != nil && !gorm.IsRecordNotFoundError(err):
		return User{}, err
	case !user.HasPassword(vm.Password):
		return User{}, fmt.Errorf("Wrong email or password")
	case user.Status == NOTACTIVE:
		return User{}, fmt.Errorf("Account requires activation")
	case user.Status == SUSPENDED:
		return User{}, fmt.Errorf("Account suspended")
	}

	if err := user.CreateJWTToken(); err != nil {
		return User{}, err
	}
	return user, nil
}

//Activate activates user
func (tr usersRepository) Activate(vm ActivateVM) (User, error) {
	if len(strings.TrimSpace(vm.Token)) == 0 {
		return User{}, fmt.Errorf("Wrong activation token")
	}
	user := User{}
	err := db.Where("token = ?", vm.Token).First(&user).Error
	switch {
	case err != nil && gorm.IsRecordNotFoundError(err):
		return User{}, fmt.Errorf("Wrong activation token")
	case err != nil && !gorm.IsRecordNotFoundError(err):
		return User{}, err
	case user.Status == SUSPENDED:
		return User{}, fmt.Errorf("Account suspended")
	}
	//update user record
	user.Status = ACTIVE
	user.Token = ""
	if err := db.Save(&user).Error; err != nil {
		return User{}, err
	}

	if err := user.CreateJWTToken(); err != nil {
		return User{}, err
	}
	return user, nil
}

//Register registers a new user
func (tr usersRepository) Register(vm RegisterVM) (User, error) {
	user := User{}
	err := db.Where("email = ?", helpers.NormalizeEmail(vm.Email)).First(&user).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return User{}, err
	}
	if user.ID != 0 && user.Status != NOTACTIVE {
		return User{}, fmt.Errorf("This email already taken")
	}

	user.Name = vm.Name
	user.Email = vm.Email
	user.PasswordHash = helpers.CreatePasswordHash(vm.Password)
	user.UserGroupID = USER
	user.Status = ACTIVE
	user.Token = ""

	//create new or update inactive account
	if err := db.Save(&user).Error; err != nil {
		return User{}, err
	}

	if err := user.CreateJWTToken(); err != nil {
		return User{}, err
	}
	return user, nil
}

//Forgot creates a secure token in case user forgor password
func (tr usersRepository) Forgot(vm ForgotVM) (User, error) {
	user := User{}
	err := db.Where("email = ?", strings.ToLower(vm.Email)).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, fmt.Errorf("User not found")
		}
		return User{}, err
	}

	user.Token = helpers.CreateSecureToken()
	if err := db.Save(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

//ResetPassword resets user password
func (tr usersRepository) ResetPassword(vm ResetVM) (User, error) {
	if len(strings.TrimSpace(vm.Token)) == 0 {
		return User{}, fmt.Errorf("Wrong token")
	}
	user := User{}
	err := db.Where("token = ?", vm.Token).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, fmt.Errorf("User not found")
		}
		return User{}, err
	}

	user.Token = ""
	user.PasswordHash = helpers.CreatePasswordHash(vm.Password)
	if err := db.Save(&user).Error; err != nil {
		return User{}, err
	}
	if err := user.CreateJWTToken(); err != nil {
		return User{}, err
	}
	return user, nil
}

//Summary returns summary info for a dashboard
func (tr usersRepository) Summary() (UsersSummaryVM, error) {
	vm := UsersSummaryVM{}
	if err := db.Model(User{}).Count(&vm.Count).Error; err != nil {
		return UsersSummaryVM{}, err
	}
	return vm, nil
}
