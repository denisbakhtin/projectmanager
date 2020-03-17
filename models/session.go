package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//Session represents a work session, where task logs (spent time) are commited
//once the task log is commited and session saved, it's is not shown in spent report
type Session struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Contents  string    `json:"contents" valid:"length(0|5000)"`
	UserID    uint64    `json:"user_id" valid:"-"`
	User      User      `json:"user" gorm:"save_associations:false" valid:"-"`
	TaskLogs  []TaskLog `json:"task_logs" gorm:"save_associations:true" valid:"-"`
}

// BeforeCreate - gorm hook, fired before record creation
func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	//check all task logs belong to the current user
	ids := make([]uint64, len(s.TaskLogs))
	for i := 0; i < len(s.TaskLogs); i++ {
		ids[i] = s.TaskLogs[i].ID
	}
	count := 0
	if tx.Where("user_id != ? and id in(?)", s.UserID, ids).Model(TaskLog{}).Count(&count); count > 0 {
		return errors.New("Some task logs do not belong to this user")
	}
	return
}

// BeforeDelete - gorm hook, validate removal here & delete other related records (no cascade tag atm)
func (s *Session) BeforeDelete(tx *gorm.DB) (err error) {
	if tx.Model(s).Association("TaskLogs").Count() > 0 {
		return fmt.Errorf("Session is not empty")
	}
	return
}

//NewSessionVM is a view model for a new session
type NewSessionVM struct {
	TaskLogs []TaskLog `json:"task_logs"`
}
