package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestCategoriesGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/categories")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/categories", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	var categories []models.Category
	assert.Nil(t, json.Unmarshal(body, &categories))
	assert.True(t, len(categories) > 0)
}

func TestCategoryGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/categories/1")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/categories/1", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	category := models.Category{}
	assert.Nil(t, json.Unmarshal(body, &category))
	assert.NotZero(t, category.ID)
}

func TestCategoriesPost(t *testing.T) {
	cat := models.Category{
		Name: "New Category",
	}
	resp, err := jsonPost(server.URL+"/api/categories", cat)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	categories, err := models.CategoriesDB.GetAll(authenticatedUser.ID)
	assert.Nil(t, err)
	count := len(categories)
	assert.True(t, count > 0)

	resp, err = jsonPostAuth(server.URL+"/api/categories", cat, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	categories, err = models.CategoriesDB.GetAll(authenticatedUser.ID)
	assert.Nil(t, err)
	assert.True(t, count+1 == len(categories))
}

func TestCategoryPut(t *testing.T) {
	category, err := models.CategoriesDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, category.ID)

	category.Name = "New Name"
	resp, err := jsonPut(server.URL+"/api/categories/9", category)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPutAuth(server.URL+"/api/categories/9", category, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	category, err = models.CategoriesDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, category.ID)
	assert.Equal(t, "New Name", category.Name)
}

func TestCategoryDelete(t *testing.T) {
	category, err := models.CategoriesDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, category.ID)

	resp, err := jsonDelete(server.URL + "/api/categories/9")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonDeleteAuth(server.URL+"/api/categories/9", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	category, err = models.CategoriesDB.Get(authenticatedUser.ID, "9")
	assert.NotNil(t, err)
}

func TestCategorySummary(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/categories_summary")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/categories_summary", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	vm := models.CategoriesSummaryVM{}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(bytes, &vm)
	assert.Nil(t, err)
	assert.Greater(t, vm.Count, 0)
}
