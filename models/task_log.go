package models

import "time"

//TaskLog represents a task workflow log
type TaskLog struct {
	ID            uint           `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     *time.Time     `sql:"index" json:"-"`
	TaskID        uint           `json:"task_id" gorm:"index"`
	Hours         float64        `json:"hours"`
	IsBillable    bool           `json:"is_billable"`
	Description   string         `json:"description" valid:"length(0|100000)"`
	ProjectUserID uint           `json:"project_user_id" gorm:"index"`
	AttachedFiles []AttachedFile `json:"files" gorm:"polymorphic:Owner" valid:"-"`
}
