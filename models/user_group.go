package models

import (
	"errors"
	"time"
)

//UserGroup represents a row from user_groups table
type UserGroup struct {
	ID         uint64     `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"-"`
	Name       string     `json:"name" valid:"required,length(1|100)"`
	Persistent bool       `json:"persistent"`
	Users      []User     `json:"users" gorm:"save_associations:false" valid:"-"`
}

//Mandatory user group IDs, adjust them accordingly if you need
const (
	ADMIN  = 1
	EDITOR = 2
	USER   = 3
)

//BeforeDelete gorm hook
func (ug *UserGroup) BeforeDelete() (err error) {
	if ug.Persistent {
		err = errors.New("Can't remove a persistent group")
	}
	count := 0
	DB.Model(&User{}).Where("user_group_id = ?", ug.ID).Count(&count)
	if count > 0 {
		err = errors.New("There are users in this group")
	}
	return
}
