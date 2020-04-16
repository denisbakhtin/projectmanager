package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestSessionsGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/sessions")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/sessions", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	var sessions []models.Session
	assert.Nil(t, json.Unmarshal(body, &sessions))
	assert.True(t, len(sessions) > 0)
}

func TestSessionGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/sessions/1")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/sessions/1", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	session := models.Session{}
	assert.Nil(t, json.Unmarshal(body, &session))
	assert.True(t, session.ID == 1)
}

func TestSessionsPost(t *testing.T) {
	cat := models.Session{
		Contents: "New Session",
	}
	resp, err := jsonPost(server.URL+"/api/sessions", cat)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	sessions, err := models.SessionsDB.GetAll(authenticatedUser.ID)
	assert.Nil(t, err)
	count := len(sessions)
	assert.True(t, count > 0)

	resp, err = jsonPostAuth(server.URL+"/api/sessions", cat, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	sessions, err = models.SessionsDB.GetAll(authenticatedUser.ID)
	assert.Nil(t, err)
	assert.True(t, count+1 == len(sessions))
}

func TestSessionDelete(t *testing.T) {
	session, err := models.SessionsDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, session.ID)

	resp, err := jsonDelete(server.URL + "/api/sessions/9")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonDeleteAuth(server.URL+"/api/sessions/9", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	session, err = models.SessionsDB.Get(authenticatedUser.ID, "9")
	assert.NotNil(t, err)
}

func TestSessionSummary(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/sessions_summary")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/sessions_summary", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	vm := models.SessionsSummaryVM{}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(bytes, &vm)
	assert.Nil(t, err)
	assert.Greater(t, vm.Count, 0)
}
