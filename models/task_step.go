package models

import (
	"errors"
	"time"
)

//TaskStep represents a task step
type TaskStep struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name      string     `json:"name" valid:"required,length(1|200)"`
	IsFinal   bool       `json:"is_final"`
	Order     uint       `json:"order,string,omitempty"` //order
}

//BeforeDelete gorm hook
func (ts *TaskStep) BeforeDelete() (err error) {
	count := 0
	DB.Model(&Task{}).Where("task_step_id = ?", ts.ID).Count(&count)
	if count > 0 {
		err = errors.New("There are tasks with this step")
	}
	return
}
