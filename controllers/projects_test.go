package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestProjectsGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/projects")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/projects", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	var projects []models.Project
	assert.Nil(t, json.Unmarshal(body, &projects))
	assert.True(t, len(projects) > 0)
}

func TestProjectGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/projects/1")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/projects/1", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	project := models.Project{}
	assert.Nil(t, json.Unmarshal(body, &project))
	assert.NotZero(t, project.ID)
}

func TestProjectsPost(t *testing.T) {
	cat := models.Project{
		Name: "New Project",
	}
	resp, err := jsonPost(server.URL+"/api/projects", cat)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	projects, err := models.ProjectsDB.GetAll(authenticatedUser.ID)
	assert.Nil(t, err)
	count := len(projects)
	assert.True(t, count > 0)

	resp, err = jsonPostAuth(server.URL+"/api/projects", cat, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	projects, err = models.ProjectsDB.GetAll(authenticatedUser.ID)
	assert.Nil(t, err)
	assert.True(t, count+1 == len(projects))
}

func TestProjectPut(t *testing.T) {
	project, err := models.ProjectsDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, project.ID)

	project.Name = "New Name"
	resp, err := jsonPut(server.URL+"/api/projects/9", project)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPutAuth(server.URL+"/api/projects/9", project, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	project, err = models.ProjectsDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, project.ID)
	assert.Equal(t, "New Name", project.Name)
}

func TestProjectDelete(t *testing.T) {
	project, err := models.ProjectsDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, project.ID)

	resp, err := jsonDelete(server.URL + "/api/projects/9")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonDeleteAuth(server.URL+"/api/projects/9", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	project, err = models.ProjectsDB.Get(authenticatedUser.ID, "9")
	assert.NotNil(t, err)
}

func TestProjectSummary(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/projects_summary")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/projects_summary", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	vm := models.ProjectsSummaryVM{}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(bytes, &vm)
	assert.Nil(t, err)
	assert.Greater(t, vm.Count, 0)
}
