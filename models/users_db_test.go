package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/stretchr/testify/assert"
)

func TestUsersRepositoryGetAll(t *testing.T) {
	getOrCreateUser()
	list, err := UsersDB.GetAll()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 1)
	for _, u := range list {
		assert.NotZero(t, u.ID)
	}
}

func TestUsersRepositoryGet(t *testing.T) {
	user := getOrCreateUser()
	p, err := UsersDB.Get(user.ID)
	assert.Nil(t, err)
	assert.NotZero(t, p.ID)
}

func TestUsersRepositoryUpdateStatus(t *testing.T) {
	user := createUser("")
	assert.Equal(t, user.Status, ACTIVE)
	user.Status = SUSPENDED
	user, err := UsersDB.UpdateStatus(user)
	assert.Nil(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, user.Status, SUSPENDED)
}

func TestUsersRepositoryUpdateAccount(t *testing.T) {
	user := createUser("123")
	vm := AccountVM{
		Name:            fmt.Sprintf("Name %d", time.Now().Nanosecond()),
		CurrentPassword: "1234",
		NewPassword:     "12345",
	}
	_, err := UsersDB.UpdateAccount(vm, user)
	assert.NotNil(t, err)

	vm.CurrentPassword = "123"
	u, err := UsersDB.UpdateAccount(vm, user)
	assert.Nil(t, err)
	assert.True(t, u.HasPassword("12345"))
	assert.Equal(t, u.Name, vm.Name)

	vm.Name = fmt.Sprintf("New name %d", time.Now().Nanosecond())
	vm.CurrentPassword = ""
	vm.NewPassword = "123456"
	u, err = UsersDB.UpdateAccount(vm, u)
	assert.Nil(t, err)
	assert.Equal(t, u.Name, vm.Name)
	assert.True(t, u.HasPassword("12345"))
}

func TestUsersRepositoryLogin(t *testing.T) {
	user := createUser("123")
	vm := LoginVM{
		Email:    "non.existent@email.com",
		Password: "12345",
	}

	_, err := UsersDB.Login(vm)
	assert.NotNil(t, err)

	vm.Email = user.Email
	_, err = UsersDB.Login(vm)
	assert.NotNil(t, err)

	vm.Password = "123"
	u, err := UsersDB.Login(vm)
	assert.Nil(t, err)
	assert.Equal(t, u.Name, user.Name)
	assert.Equal(t, u.Email, user.Email)
	assert.True(t, len(u.JWTToken) > 0)
}

func TestUsersRepositoryActivate(t *testing.T) {
	user := createUser("123")
	user.Status = NOTACTIVE
	user.Token = helpers.CreateSecureToken()
	err := DB.Save(&user).Error
	assert.Nil(t, err)

	vm := ActivateVM{
		Token: user.Token + "1",
	}
	_, err = UsersDB.Activate(vm)
	assert.NotNil(t, err)

	vm.Token = user.Token
	u, err := UsersDB.Activate(vm)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, user.ID)
	assert.True(t, len(u.Token) == 0)
	assert.True(t, len(u.JWTToken) > 0)
	assert.Equal(t, u.Status, ACTIVE)
}

func TestUsersRepositoryRegister(t *testing.T) {
	vm := RegisterVM{
		Name:     fmt.Sprintf("Name %d", time.Now().Nanosecond()),
		Email:    fmt.Sprintf("%d@email.com", time.Now().Nanosecond()),
		Password: "12345",
	}
	user, err := UsersDB.Register(vm)
	assert.Nil(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, vm.Name, user.Name)
	assert.Equal(t, vm.Email, user.Email)
	assert.True(t, user.HasPassword(vm.Password))
	assert.Equal(t, user.Status, ACTIVE)
	assert.True(t, len(user.JWTToken) > 0)
}

func TestUsersRepositorySummary(t *testing.T) {
	//atleast one user exists
	getOrCreateUser()
	vm, err := UsersDB.Summary()
	assert.Nil(t, err)
	assert.NotEmpty(t, vm)
	assert.Greater(t, vm.Count, 0)
}
