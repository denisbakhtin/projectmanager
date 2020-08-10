package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestTaskLogsPost(t *testing.T) {
	tl := models.TaskLog{
		Minutes: 1,
		TaskID:  1, //TODO: substitute
	}
	resp, err := jsonPost(server.URL+"/api/task_logs", tl)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPostAuth(server.URL+"/api/task_logs", tl, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTaskLogsPut(t *testing.T) {
	tl := models.TaskLog{
		ID:      111,
		Minutes: 2,
		TaskID:  1, //TODO: substitute
	}
	resp, err := jsonPut(server.URL+"/api/task_logs/111", tl)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPutAuth(server.URL+"/api/task_logs/111", tl, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTaskLogsLatestGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/task_logs_latest")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/task_logs_latest", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var logs []models.TaskLog
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(bytes, &logs)
	assert.Nil(t, err)
	assert.Greater(t, len(logs), 0)
	assert.LessOrEqual(t, len(logs), 5)
}
