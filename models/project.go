package models

import (
	"time"
)

//Project represents a record from projects table
type Project struct {
	ID            uint           `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     *time.Time     `sql:"index" json:"-"`
	Name          string         `json:"name" valid:"required,length(1|1500)"`
	Description   string         `json:"description" valid:"length(0|100000)"`
	OwnerID       uint           `json:"owner_id" valid:"required"`
	StartDate     *time.Time     `json:"start_date"`
	EndDate       *time.Time     `json:"end_date"`
	StatusID      uint           `json:"status_id,string,omitempty" valid:"required"`
	AttachedFiles []AttachedFile `json:"files" gorm:"polymorphic:Owner" valid:"-"`
	ProjectUsers  []ProjectUser  `json:"project_users" valid:"-"`
	Tasks         []Task         `json:"tasks" gorm:"save_associations:false" valid:"-"`
	Owner         User           `json:"owner" gorm:"save_associations:false" valid:"-"`
	Status        Status         `json:"status" gorm:"save_associations:false" valid:"-"`
}
