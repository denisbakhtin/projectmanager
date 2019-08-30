package models

import "time"

//Task represents a row from tasks table
type Task struct {
	ID            uint64         `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     *time.Time     `sql:"index" json:"-"`
	Name          string         `json:"name" valid:"required,length(1|1500)"`
	Description   string         `json:"description" valid:"length(0|100000)"`
	ProjectID     uint64         `json:"project_id" gorm:"index" valid:"required"`
	ProjectUserID uint64         `json:"project_user_id" gorm:"index" valid:"required"`
	TaskStepID    uint64         `json:"task_step_id" gorm:"index" valid:"required"`
	AttachedFiles []AttachedFile `json:"files" gorm:"polymorphic:Owner" valid:"-"`
	Project       Project        `json:"project" gorm:"save_associations:false" valid:"-"`
	ProjectUser   ProjectUser    `json:"project_user" gorm:"save_associations:false" valid:"-"`
	TaskStep      TaskStep       `json:"task_step" gorm:"save_associations:false" valid:"-"`
}
