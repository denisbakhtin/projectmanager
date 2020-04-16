package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskLogsRepositoryCreate(t *testing.T) {
	minutes := uint64(time.Now().Nanosecond())
	task := getOrCreateTask()
	tl := TaskLog{
		Minutes: minutes,
		TaskID:  task.ID,
	}
	var logs []TaskLog
	err := DB.Find(&logs).Error
	assert.Nil(t, err)
	count := len(logs)
	l, err := TaskLogsDB.Create(userID, tl)
	assert.Nil(t, err)
	assert.NotZero(t, l.ID)
	assert.Equal(t, l.Minutes, minutes)
	assert.Equal(t, l.UserID, userID)
	assert.Zero(t, l.SessionID)

	err = DB.First(&l, l.ID).Error
	assert.Nil(t, err)
	assert.NotZero(t, l.ID)
	assert.Equal(t, l.Minutes, minutes)

	DB.Find(&logs)
	assert.Equal(t, count+1, len(logs), "The number of task logs should have been increased by 1")
}

func TestTaskLogsRepositoryUpdate(t *testing.T) {
	minutes := uint64(time.Now().Nanosecond())
	tl := TaskLog{}
	DB.Where("user_id = ? and session_id != 0", userID).First(&tl)
	assert.NotZero(t, tl.ID)
	tl.Minutes = minutes
	_, err := TaskLogsDB.Update(userID, tl)
	assert.NotNil(t, err)

	tl = getOrCreateTaskLog()
	assert.NotZero(t, tl.ID)
	assert.NotEqual(t, tl.Minutes, minutes)

	tl.Minutes = minutes
	tl, err = TaskLogsDB.Update(userID, tl)
	assert.Nil(t, err)
	assert.Equal(t, tl.UserID, userID)
	assert.Equal(t, tl.Minutes, minutes)
}
