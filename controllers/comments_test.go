package controllers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/stretchr/testify/assert"
)

func TestCommentsGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/comments")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/comments", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	var comments []models.Comment
	assert.Nil(t, json.Unmarshal(body, &comments))
	assert.True(t, len(comments) > 0)
}

func TestCommentGet(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/comments/1")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonGetAuth(server.URL+"/api/comments/1", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	comment := models.Comment{}
	assert.Nil(t, json.Unmarshal(body, &comment))
	assert.True(t, comment.ID == 1)
}

func TestCommentsPost(t *testing.T) {
	cat := models.Comment{
		Contents: "New Comment",
		TaskID:   1,
	}
	resp, err := jsonPost(server.URL+"/api/comments", cat)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	comments, err := models.CommentsDB.GetAll(authenticatedUser.ID, 1)
	assert.Nil(t, err)
	count := len(comments)
	assert.True(t, count > 0)

	resp, err = jsonPostAuth(server.URL+"/api/comments", cat, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	comments, err = models.CommentsDB.GetAll(authenticatedUser.ID, 1)
	assert.Nil(t, err)
	assert.True(t, count+1 == len(comments))
}

func TestCommentPut(t *testing.T) {
	comment, err := models.CommentsDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, comment.ID)

	comment.Contents = "New Content"
	resp, err := jsonPut(server.URL+"/api/comments/9", comment)
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonPutAuth(server.URL+"/api/comments/9", comment, authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	comment, err = models.CommentsDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, comment.ID)
	assert.Equal(t, "New Content", comment.Contents)
}

func TestCommentDelete(t *testing.T) {
	comment, err := models.CommentsDB.Get(authenticatedUser.ID, "9")
	assert.Nil(t, err)
	assert.NotZero(t, comment.ID)

	resp, err := jsonDelete(server.URL + "/api/comments/9")
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	resp, err = jsonDeleteAuth(server.URL+"/api/comments/9", authenticatedUser.JWTToken)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	comment, err = models.CommentsDB.Get(authenticatedUser.ID, "9")
	assert.NotNil(t, err)
}
