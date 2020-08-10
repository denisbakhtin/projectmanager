package models

import "fmt"

//SessionsDB is a sessions db repository
var SessionsDB SessionsRepository

func init() {
	SessionsDB = &sessionsRepository{}
}

//SessionsRepository is a repository of sessions
type SessionsRepository interface {
	GetAll(userID uint64) ([]Session, error)
	Get(userID uint64, id interface{}) (Session, error)
	NewGet(userID uint64) (Session, error)
	Create(userID uint64, session Session) (Session, error)
	Delete(userID uint64, id interface{}) error
	Summary(userID uint64) (SessionsSummaryVM, error)
}

type sessionsRepository struct{}

//GetAll returns all sessions owned by specified user
func (sr *sessionsRepository) GetAll(userID uint64) ([]Session, error) {
	var sessions []Session
	err := db.Where("user_id = ?", userID).Preload("TaskLogs").Find(&sessions).Error
	return sessions, err
}

//Get fetches a session by its id
func (sr *sessionsRepository) Get(userID uint64, id interface{}) (Session, error) {
	session := Session{}
	err := db.Where("user_id = ?", userID).
		Preload("TaskLogs").Preload("TaskLogs.Task").Preload("TaskLogs.Task.Project").
		First(&session, id).Error
	return session, err
}

//NewGet gets a view models for a new session
func (sr *sessionsRepository) NewGet(userID uint64) (Session, error) {
	var session Session
	err := db.Where("user_id = ? and minutes > 0 and session_id = 0", userID).
		Preload("Task").Preload("Task.Project").Find(&session.TaskLogs).Error
	return session, err
}

//Create sreates a new session in db
func (sr *sessionsRepository) Create(userID uint64, session Session) (Session, error) {
	session.UserID = userID
	err := db.Create(&session).Error
	return session, err
}

//Delete removes a session from db
func (sr *sessionsRepository) Delete(userID uint64, id interface{}) error {
	session := Session{}
	err := db.Preload("TaskLogs").Where("user_id = ?", userID).First(&session, id).Error
	if err != nil {
		return err
	}
	if len(session.TaskLogs) > 0 {
		return fmt.Errorf("Can not remove non-empty session")
	}
	if err := db.Delete(&session).Error; err != nil {
		return err
	}
	return nil
}

//Summary returns summary info for a dashboard
func (sr *sessionsRepository) Summary(userID uint64) (SessionsSummaryVM, error) {
	vm := SessionsSummaryVM{}
	err := db.Model(Session{}).Where("user_id = ?", userID).Count(&vm.Count).Error
	return vm, err
}
