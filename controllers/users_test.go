package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestUsersGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/users")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/users", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/users", authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	var users []models.User
	assert.Nil(t, json.Unmarshal(body, &users))
	assert.True(t, len(users) > 0)
}

func TestUserGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/users/1")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/users/1", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/users/1", authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	user := models.User{}
	assert.Nil(t, json.Unmarshal(body, &user))
	assert.True(t, user.ID == 1)
}

func TestUserStatusPut(t *testing.T) {
	user, err := models.UsersDB.Get("9")
	assert.Nil(t, err)
	assert.True(t, user.ID == 9)

	user.Status = models.SUSPENDED
	resp, err := jsonPut(server.URL+"/api/users/9", user)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPutAuth(server.URL+"/api/users/9", user, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPutAuth(server.URL+"/api/users/9", user, authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	user, err = models.UsersDB.Get("9")
	assert.Nil(t, err)
	assert.True(t, user.ID == 9)
	assert.Equal(t, models.SUSPENDED, user.Status)
}

func TestUserSummary(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/users_summary")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/users_summary", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/users_summary", authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	vm := models.UsersSummaryVM{}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(bytes, &vm)
	assert.Nil(t, err)
	assert.Greater(t, vm.Count, 0)
}
