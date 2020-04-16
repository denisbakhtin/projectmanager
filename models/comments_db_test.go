package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestCommentsRepositoryGetAll(t *testing.T) {
	task := getOrCreateTask()
	getOrCreateComment(task.ID)
	list, err := CommentsDB.GetAll(userID, task.ID)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 1)
	for _, com := range list {
		assert.Equal(t, com.UserID, userID, "Comment should belong to the specified user")
	}
}

func TestCommentsRepositoryGet(t *testing.T) {
	task := getOrCreateTask()
	comment := getOrCreateComment(task.ID)
	com, err := CommentsDB.Get(userID, comment.ID)
	assert.Nil(t, err)
	assert.NotZero(t, com.ID)
	assert.Equal(t, com.UserID, userID)
}

func TestCommentsRepositoryCreate(t *testing.T) {
	contents := fmt.Sprintf("Comment-%d", time.Now().Nanosecond())
	task := getOrCreateTask()
	comment := Comment{
		Contents: contents,
		TaskID:   task.ID,
	}

	all, _ := CommentsDB.GetAll(userID, task.ID)
	count := len(all)
	cat, err := CommentsDB.Create(userID, comment)
	assert.Nil(t, err)
	assert.NotZero(t, cat.ID)
	assert.Equal(t, cat.Contents, contents)

	cat, err = CommentsDB.Get(userID, cat.ID)
	assert.Nil(t, err)
	assert.NotZero(t, cat.ID)
	assert.Equal(t, cat.Contents, contents)

	all, _ = CommentsDB.GetAll(userID, task.ID)
	assert.Equal(t, count+1, len(all), "The number of comments should have been increased by 1")
}

func TestCommentsRepositoryUpdate(t *testing.T) {
	contents := fmt.Sprintf("Comment-%d", time.Now().Nanosecond())
	comment := Comment{}
	DB.Where("user_id = ?", userID).First(&comment)
	assert.NotZero(t, comment.ID)
	assert.NotEqual(t, comment.Contents, contents)

	comment.Contents = contents
	cat, err := CommentsDB.Update(userID, comment)
	assert.Nil(t, err)
	assert.Equal(t, cat.UserID, userID)
	assert.Equal(t, cat.Contents, contents)
}

func TestCommentsRepositoryDelete(t *testing.T) {
	err := CommentsDB.Delete(userID, 11111111111)
	assert.NotNil(t, err)
	assert.True(t, gorm.IsRecordNotFoundError(err), "This comment should not exist")

	task := getOrCreateTask()
	comment := getOrCreateComment(task.ID)
	assert.NotZero(t, comment.ID)

	err = CommentsDB.Delete(userID+1, comment.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Should check comment owner")

	err = CommentsDB.Delete(userID, comment.ID)
	assert.Nil(t, err)
	_, err = CommentsDB.Get(userID, comment.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Comment should have been removed")
}
