package mocks

import (
	"fmt"
	"strconv"

	"github.com/denisbakhtin/projectmanager/models"
)

//SessionsDBMock is a SessionsDB repository mock
type SessionsDBMock struct {
	Sessions []models.Session
}

func (cr *SessionsDBMock) GetAll(userID uint64) ([]models.Session, error) {
	return cr.Sessions, nil
}

func (cr *SessionsDBMock) Get(userID uint64, id interface{}) (models.Session, error) {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for _, c := range cr.Sessions {
		if c.ID == idi {
			return c, nil
		}
	}
	return models.Session{}, fmt.Errorf("Session not found")
}

//Create creates a new session in db
func (cr *SessionsDBMock) Create(userID uint64, session models.Session) (models.Session, error) {
	session.UserID = userID
	cr.Sessions = append(cr.Sessions, session)
	return session, nil
}

//NewGet gets a view models for a new session
func (cr *SessionsDBMock) NewGet(userID uint64) (models.Session, error) {
	return models.Session{}, nil
}

//Delete removes a session from db
func (cr *SessionsDBMock) Delete(userID uint64, id interface{}) error {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range cr.Sessions {
		if cr.Sessions[i].ID == idi && cr.Sessions[i].UserID == userID {
			cr.Sessions[i] = cr.Sessions[len(cr.Sessions)-1] // Copy last element to index i.
			cr.Sessions = cr.Sessions[:len(cr.Sessions)-1]   // Truncate slice.
			return nil
		}
	}
	return fmt.Errorf("Session not found")
}

//Summary returs session summary
func (cr *SessionsDBMock) Summary(userID uint64) (models.SessionsSummaryVM, error) {
	return models.SessionsSummaryVM{Count: len(cr.Sessions)}, nil
}
