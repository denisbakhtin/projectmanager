package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestAccountGet(t *testing.T) {
	assert.True(t, len(authenticatedUser.JWTToken) > 0)
	resp, err := jsonGetAuth(server.URL+"/api/account", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	user := models.User{}
	assert.Nil(t, json.Unmarshal(body, &user))
	assert.Equal(t, authenticatedUser.Email, user.Email)
}

func TestAccountPut(t *testing.T) {
	u, err := models.UsersDB.GetByEmail("4@email.com")
	assert.Nil(t, err)
	assert.Equal(t, u.Email, "4@email.com")
	u.CreateJWTToken()
	assert.True(t, len(u.JWTToken) > 0)

	vm := models.AccountVM{
		Name:            "NewName",
		CurrentPassword: "4",
		NewPassword:     "444",
	}
	resp, err := jsonPutAuth(server.URL+"/api/account", vm, u.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	u, err = models.UsersDB.GetByEmail("4@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "NewName", u.Name)
	assert.True(t, u.HasPassword("444"))

	vm.CurrentPassword = "wrong password"
	resp, err = jsonPutAuth(server.URL+"/api/account", vm, u.JWTToken)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, resp.StatusCode, 400)
	u, err = models.UsersDB.GetByEmail("4@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "NewName", u.Name)
	assert.True(t, u.HasPassword("444"))
}
