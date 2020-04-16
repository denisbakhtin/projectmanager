package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

//Mandatory user group IDs, adjust them accordingly if you need
const (
	ADMIN  = 1
	EDITOR = 2
	USER   = 3
)

//UserGroup represents a row from user_groups table
type UserGroup struct {
	ID         uint64    `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Name       string    `json:"name" valid:"required,length(1|100)"`
	Persistent bool      `json:"persistent"`
	Users      []User    `json:"users" gorm:"save_associations:false" valid:"-"`
}

//BeforeDelete gorm hook
func (ug *UserGroup) BeforeDelete(tx *gorm.DB) (err error) {
	if ug.Persistent {
		err = errors.New("Can't remove a persistent group")
	}

	if tx.Model(ug).Association("Users").Count() > 0 {
		err = errors.New("There are users in this group")
	}
	return
}
