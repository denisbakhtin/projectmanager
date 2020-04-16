package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestPagesGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/pages")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)
	resp, err = jsonGetAuth(server.URL+"/api/pages", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)
	resp, err = jsonGetAuth(server.URL+"/api/pages", authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	var pages []models.Page
	assert.Nil(t, json.Unmarshal(body, &pages))
	assert.True(t, len(pages) > 0)
}

func TestPageGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/pages/1")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/pages/1", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/pages/1", authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	page := models.Page{}
	assert.Nil(t, json.Unmarshal(body, &page))
	assert.True(t, page.ID == 1)
}

func TestPagesPost(t *testing.T) {
	p := models.Page{
		Name: "New Page",
	}
	resp, err := jsonPost(server.URL+"/api/pages", p)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)
	resp, err = jsonPostAuth(server.URL+"/api/pages", p, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	pages, err := models.PagesDB.GetAll()
	assert.Nil(t, err)
	count := len(pages)
	assert.True(t, count > 0)

	resp, err = jsonPostAuth(server.URL+"/api/pages", p, authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	pages, err = models.PagesDB.GetAll()
	assert.Nil(t, err)
	assert.True(t, count+1 == len(pages))
}

func TestPagePut(t *testing.T) {
	page, err := models.PagesDB.Get("9")
	assert.Nil(t, err)
	assert.NotZero(t, page.ID)

	page.Name = "New Name"
	resp, err := jsonPut(server.URL+"/api/pages/9", page)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)
	resp, err = jsonPutAuth(server.URL+"/api/pages/9", page, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPutAuth(server.URL+"/api/pages/9", page, authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	page, err = models.PagesDB.Get("9")
	assert.Nil(t, err)
	assert.NotZero(t, page.ID)
	assert.Equal(t, "New Name", page.Name)
}

func TestPageDelete(t *testing.T) {
	page, err := models.PagesDB.Get("9")
	assert.Nil(t, err)
	assert.NotZero(t, page.ID)

	resp, err := jsonDelete(server.URL + "/api/pages/9")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)
	resp, err = jsonDeleteAuth(server.URL+"/api/pages/9", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonDeleteAuth(server.URL+"/api/pages/9", authenticatedAdmin.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	page, err = models.PagesDB.Get("9")
	assert.NotNil(t, err)
}

func TestPageGetHTML(t *testing.T) {
	resp, err := jsonGet(server.URL + "/pages/1")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	resp, err = jsonGet(server.URL + "/pages/444")
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, resp.StatusCode, 400)
}
