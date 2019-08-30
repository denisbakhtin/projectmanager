package models

import "time"

//ProjectUser represents a record in project_users table
type ProjectUser struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	UserID    uint64     `json:"user_id" gorm:"index:idx" valid:"required"`
	RoleID    uint64     `json:"role_id,string,omitempty" gorm:"index:idx" valid:"required"`
	ProjectID uint64     `json:"project_id" gorm:"index:idx" valid:"required"`
	User      User       `json:"user" gorm:"save_associations:false" valid:"-"`
	Role      Role       `json:"role" gorm:"save_associations:false" valid:"-"`
	Project   Project    `json:"project" gorm:"save_associations:false" valid:"-"`
}
