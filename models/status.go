package models

import "time"

//Status represents a project status
type Status struct {
	ID          uint64     `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"-"`
	Name        string     `json:"name" valid:"required,length(1|100)"`
	Description string     `json:"description" valid:"length(0|1500)"`
	Ord         uint64     `json:"ord,string,omitempty"` //order
}
