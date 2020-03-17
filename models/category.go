package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/jinzhu/gorm"
)

//Category represents a project or task category (n:1 atm)
type Category struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" valid:"required,length(1|1500)"`
	Tasks     []Task    `json:"tasks" gorm:"save_associations:false" valid:"-"`
	Projects  []Project `json:"projects" gorm:"save_associations:false" valid:"-"`
	UserID    uint64    `json:"user_id" valid:"-"`
	User      User      `json:"user" gorm:"save_associations:false" valid:"-"`
}

// BeforeUpdate - gorm hook, fired before record update
func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	//check if original user_id is not being changed
	cat := Category{}
	if tx.Where("id = ? and user_id = ?", c.ID, c.UserID).First(&cat); cat.ID == 0 {
		return errors.New(helpers.NotFoundOrOwned("Category"))
	}
	return
}

// BeforeDelete - gorm hook, validate removal here
func (c *Category) BeforeDelete(tx *gorm.DB) (err error) {
	//check existence of associated records and prohibit if any
	if tx.Model(c).Association("Projects").Count() > 0 {
		return fmt.Errorf("Category has associated projects with it")
	}
	if tx.Model(c).Association("Tasks").Count() > 0 {
		return fmt.Errorf("Category has associated tasks with it")
	}
	return
}
