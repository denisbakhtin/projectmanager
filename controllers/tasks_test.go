package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestTasksGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/tasks")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/tasks", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	var tasks []models.Task
	assert.Nil(t, json.Unmarshal(body, &tasks))
	assert.True(t, len(tasks) > 0)
}

func TestTaskGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/tasks/1")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/tasks/1", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	task := models.Task{}
	assert.Nil(t, json.Unmarshal(body, &task))
	assert.True(t, task.ID == 1)
}

func TestTasksPost(t *testing.T) {
	cat := models.Task{
		Name: "New Task",
	}
	resp, err := jsonPost(server.URL+"/api/tasks", cat)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	tasks, err := models.TasksDB.GetAll(authenticatedUser.ID)
	assert.Nil(t, err)
	count := len(tasks)
	assert.True(t, count > 0)

	resp, err = jsonPostAuth(server.URL+"/api/tasks", cat, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	tasks, err = models.TasksDB.GetAll(authenticatedUser.ID)
	assert.Nil(t, err)
	assert.True(t, count+1 == len(tasks))
}

func TestTaskPut(t *testing.T) {
	task, err := models.TasksDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, task.ID)

	task.Name = "New Name"
	resp, err := jsonPut(server.URL+"/api/tasks/9", task)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPutAuth(server.URL+"/api/tasks/9", task, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	task, err = models.TasksDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, task.ID)
	assert.Equal(t, "New Name", task.Name)
}

func TestTaskDelete(t *testing.T) {
	task, err := models.TasksDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, task.ID)

	resp, err := jsonDelete(server.URL + "/api/tasks/9")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonDeleteAuth(server.URL+"/api/tasks/9", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	task, err = models.TasksDB.Get(authenticatedUser.ID, "9")
	assert.NotNil(t, err)
}

func TestTaskSummary(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/tasks_summary")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/tasks_summary", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	vm := models.TasksSummaryVM{}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(bytes, &vm)
	assert.Nil(t, err)
	assert.Greater(t, vm.Count, 0)
}
