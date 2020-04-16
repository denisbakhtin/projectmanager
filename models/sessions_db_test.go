package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestSessionsRepositoryGetAll(t *testing.T) {
	list, err := SessionsDB.GetAll(userID)
	assert.Nil(t, err, "Error should be nil")
	assert.GreaterOrEqual(t, len(list), 1)
	for _, s := range list {
		assert.Equal(t, s.UserID, userID, "Session should belong to the specified user")
	}
}

func TestSessionsRepositoryGet(t *testing.T) {
	session := getOrCreateUnrelatedSession()
	s, err := SessionsDB.Get(userID, session.ID)
	assert.Nil(t, err)
	assert.Equal(t, s.UserID, userID)
	assert.Equal(t, s.ID, session.ID)
}

func TestSessionsRepositoryCreate(t *testing.T) {
	contents := fmt.Sprintf("Session-%d", time.Now().Nanosecond())
	session := Session{
		Contents: contents,
	}
	all, _ := SessionsDB.GetAll(userID)
	count := len(all)
	s, err := SessionsDB.Create(userID, session)
	assert.Nil(t, err)
	assert.NotZero(t, s.ID)
	assert.Equal(t, s.Contents, contents)

	s, err = SessionsDB.Get(userID, s.ID)
	assert.Nil(t, err)
	assert.NotZero(t, s.ID)
	assert.Equal(t, s.Contents, contents)

	all, _ = SessionsDB.GetAll(userID)
	assert.Equal(t, count+1, len(all), "The number of sessions should have been increased by 1")
}

func TestSessionsRepositoryDelete(t *testing.T) {
	err := SessionsDB.Delete(userID, 11111111111)
	assert.NotNil(t, err)
	assert.True(t, gorm.IsRecordNotFoundError(err), "This session should not exist")

	session := getOrCreateUnrelatedSession()
	assert.NotZero(t, session.ID)

	err = SessionsDB.Delete(userID+1, session.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Should check session owner")

	err = SessionsDB.Delete(userID, session.ID)
	assert.Nil(t, err)
	_, err = SessionsDB.Get(userID, session.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Session should have been removed")

	session = getOrCreateRelatedSession()
	assert.NotZero(t, session.ID)
	err = SessionsDB.Delete(userID, session.ID)
	assert.NotNil(t, err, "Should not remove sessions with logs")
}

func TestSessionsRepositorySummary(t *testing.T) {
	//atleast one related session exists
	vm, err := SessionsDB.Summary(userID)
	assert.Nil(t, err)
	assert.NotEmpty(t, vm)
	assert.Greater(t, vm.Count, 0)
}
