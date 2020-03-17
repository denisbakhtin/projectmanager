package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/jinzhu/gorm"
)

//TaskLog represents a task workflow log
type TaskLog struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TaskID    uint64    `json:"task_id" gorm:"index" valid:"required"`
	Minutes   uint64    `json:"minutes"`
	Commited  bool      `json:"commited"`
	UserID    uint64    `json:"user_id" valid:"-"`
	User      User      `json:"user" gorm:"save_associations:false" valid:"-"`
	Task      Task      `json:"task" gorm:"save_associations:false" valid:"-"`
	SessionID uint64    `json:"session_id" gorm:"index"`
	Session   Session   `json:"session" gorm:"save_associations:false" valid:"-"`
}

// BeforeUpdate - gorm hook, fired before record update
func (t *TaskLog) BeforeUpdate(tx *gorm.DB) (err error) {
	//check if original user_id is not being changed
	tl := TaskLog{}
	if tx.Where("id = ? and user_id = ?", t.ID, t.UserID).First(&tl); tl.ID == 0 {
		return errors.New(helpers.NotFoundOrOwned("Task Log"))
	}
	return
}

// BeforeDelete - gorm hook, validate removal here & delete other related records (no cascade tag atm)
func (t *TaskLog) BeforeDelete(tx *gorm.DB) (err error) {
	//prohibit deletion
	return fmt.Errorf("Can't remove tracked task activities")
}
