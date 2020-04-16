package mocks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
)

//UsersDBMock is a UsersDB repository mock
type UsersDBMock struct {
	Users []models.User
}

func (u *UsersDBMock) GetAll() ([]models.User, error) {
	return u.Users, nil
}

func (u *UsersDBMock) Get(id interface{}) (models.User, error) {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for _, us := range u.Users {
		if us.ID == idi {
			return us, nil
		}
	}
	return models.User{}, fmt.Errorf("User not found")
}

func (u *UsersDBMock) GetByEmail(email string) (models.User, error) {
	email = helpers.NormalizeEmail(email)
	for _, us := range u.Users {
		if us.Email == email {
			return us, nil
		}
	}
	return models.User{}, fmt.Errorf("User not found")
}

func (u *UsersDBMock) UpdateStatus(user models.User) (models.User, error) {
	for i := range u.Users {
		if u.Users[i].ID == user.ID {
			u.Users[i].Status = user.Status
			return u.Users[i], nil
		}
	}
	return models.User{}, fmt.Errorf("User not found")
}

func (u *UsersDBMock) UpdateAccount(vm models.AccountVM, user models.User) (models.User, error) {
	for i := range u.Users {
		if u.Users[i].ID == user.ID {
			u2 := &u.Users[i]
			u2.Name = vm.Name
			if len(vm.CurrentPassword) > 0 && len(vm.NewPassword) > 0 {
				if !user.HasPassword(vm.CurrentPassword) {
					return models.User{}, fmt.Errorf("Wrong current password")
				}
				u2.PasswordHash = helpers.CreatePasswordHash(vm.NewPassword)
			}
			return *u2, nil
		}
	}
	return models.User{}, fmt.Errorf("User not found")
}

func (u *UsersDBMock) Login(vm models.LoginVM) (models.User, error) {
	for _, u2 := range u.Users {
		if u2.Email == helpers.NormalizeEmail(vm.Email) {
			if !u2.HasPassword(vm.Password) {
				return models.User{}, fmt.Errorf("Wrong email or password")
			}
			if u2.Status != models.ACTIVE {
				return models.User{}, fmt.Errorf("Account not active")
			}
			if err := u2.CreateJWTToken(); err != nil {
				return models.User{}, err
			}
			return u2, nil
		}
	}
	return models.User{}, fmt.Errorf("Wrong email or password")
}

func (u *UsersDBMock) Activate(vm models.ActivateVM) (models.User, error) {
	if len(strings.TrimSpace(vm.Token)) == 0 {
		return models.User{}, fmt.Errorf("Wrong token")
	}
	for i := range u.Users {
		if u.Users[i].Token == vm.Token {
			u.Users[i].Status = models.ACTIVE
			u.Users[i].Token = ""
			if err := u.Users[i].CreateJWTToken(); err != nil {
				return models.User{}, err
			}
			return u.Users[i], nil
		}
	}
	return models.User{}, fmt.Errorf("Wrong token")
}

func (u *UsersDBMock) Register(vm models.RegisterVM) (models.User, error) {
	for _, u2 := range u.Users {
		if u2.Email == helpers.NormalizeEmail(vm.Email) {
			return models.User{}, fmt.Errorf("Email already taken")
		}
	}
	u2 := models.User{
		Name:         vm.Name,
		Email:        vm.Email,
		PasswordHash: helpers.CreatePasswordHash(vm.Password),
		UserGroupID:  models.USER,
		Status:       models.ACTIVE,
		Token:        "",
	}
	if err := u2.CreateJWTToken(); err != nil {
		return models.User{}, err
	}
	u.Users = append(u.Users, u2)
	return u2, nil
}

func (u *UsersDBMock) Forgot(vm models.ForgotVM) (models.User, error) {
	for i := range u.Users {
		u2 := &u.Users[i]
		if u2.Email == helpers.NormalizeEmail(vm.Email) {
			u2.Token = helpers.CreateSecureToken()
			return *u2, nil
		}
	}
	return models.User{}, fmt.Errorf("User not found")
}

func (u *UsersDBMock) ResetPassword(vm models.ResetVM) (models.User, error) {
	for i := range u.Users {
		u2 := &u.Users[i]
		if u2.Token == vm.Token {
			u2.Token = ""
			u2.PasswordHash = helpers.CreatePasswordHash(vm.Password)
			if err := u2.CreateJWTToken(); err != nil {
				return models.User{}, err
			}
			return *u2, nil
		}
	}
	return models.User{}, fmt.Errorf("User not found")
}

func (u *UsersDBMock) Summary() (models.UsersSummaryVM, error) {
	return models.UsersSummaryVM{Count: len(u.Users)}, nil
}
