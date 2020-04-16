package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

type Token struct {
	Token string `json:"token"`
}

func TestRegister(t *testing.T) {
	vm := models.RegisterVM{
		Email:    fmt.Sprintf("%d@email.com", time.Now().Nanosecond()),
		Name:     fmt.Sprintf("Name %d", time.Now().Nanosecond()),
		Password: "12345678",
	}
	resp, err := jsonPost(server.URL+"/api/register", vm)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	token := Token{}
	assert.Nil(t, json.Unmarshal(body, &token))
	assert.True(t, len(token.Token) > 0)
	assert.True(t, emailMock.RegistrationSent)
}

func TestLogin(t *testing.T) {
	vm := models.LoginVM{
		Email:    "0@email.com",
		Password: "0",
	}
	resp, err := jsonPost(server.URL+"/api/login", vm)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	token := Token{}
	assert.Nil(t, json.Unmarshal(body, &token))
	assert.True(t, len(token.Token) > 0)

	vm.Email = "any@other.email"
	resp, err = jsonPost(server.URL+"/api/login", vm)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, resp.StatusCode, 400)

	vm.Email = "0@email.com"
	vm.Password = "any other"
	resp, err = jsonPost(server.URL+"/api/login", vm)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, resp.StatusCode, 400)
}

func TestActivate(t *testing.T) {
	vm := models.ActivateVM{
		Token: userActivationToken,
	}
	resp, err := jsonPost(server.URL+"/api/activate", vm)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	token := Token{}
	assert.Nil(t, json.Unmarshal(body, &token))
	assert.True(t, len(token.Token) > 0)

	//second time should fail
	resp, err = jsonPost(server.URL+"/api/activate", vm)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, resp.StatusCode, 400)

	vm.Token = "any other token"
	resp, err = jsonPost(server.URL+"/api/activate", vm)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, resp.StatusCode, 400)
}

func TestForgot(t *testing.T) {
	vm := models.ForgotVM{
		Email: "1@email.com",
	}

	//TODO: mock email services

	resp, err := jsonPost(server.URL+"/api/forgot", vm)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, emailMock.ResetSent)
	user, err := models.UsersDB.GetByEmail("1@email.com")
	assert.Nil(t, err)
	assert.True(t, len(user.Token) > 0)
}

func TestReset(t *testing.T) {
	vm := models.ForgotVM{
		Email: "1@email.com",
	}

	resp, err := jsonPost(server.URL+"/api/forgot", vm)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, emailMock.ResetSent)
	user, err := models.UsersDB.GetByEmail("1@email.com")
	assert.Nil(t, err)
	assert.True(t, len(user.Token) > 0)
}

func TestResetConfirmation(t *testing.T) {
	vm := models.ForgotVM{
		Email: "2@email.com",
	}

	resp, err := jsonPost(server.URL+"/api/forgot", vm)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, emailMock.ResetSent)
	user, err := models.UsersDB.GetByEmail("2@email.com")
	assert.Nil(t, err)
	assert.True(t, len(user.Token) > 0)

	vm2 := models.ResetVM{
		Token:    user.Token,
		Password: "123",
	}
	resp, err = jsonPost(server.URL+"/api/reset", vm2)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, emailMock.ResetConfirmationSent)
	user, err = models.UsersDB.GetByEmail("2@email.com")
	assert.Nil(t, err)
	assert.Equal(t, user.Email, "2@email.com")
	assert.True(t, user.HasPassword("123"))
	assert.True(t, len(user.Token) == 0)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	token := Token{}
	assert.Nil(t, json.Unmarshal(body, &token))
	assert.True(t, len(token.Token) > 0)
}
