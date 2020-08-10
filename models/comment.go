package models

import (
	"errors"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/jinzhu/gorm"
)

//Comment represents a task comment
type Comment struct {
	ID            uint64         `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Contents      string         `json:"contents" valid:"required,length(1|5000)"`
	IsSolution    bool           `json:"is_solution"`
	UserID        uint64         `json:"user_id" valid:"-"`
	User          User           `json:"user" gorm:"save_associations:false" valid:"-"`
	TaskID        uint64         `json:"task_id" gorm:"index" valid:"required"`
	Task          Task           `json:"task" gorm:"save_associations:false" valid:"-"`
	AttachedFiles []AttachedFile `json:"files" gorm:"polymorphic:Owner" valid:"-"`
	//ProjectUserID uint64         `json:"project_user_id" gorm:"index"`
	//ProjectUser   ProjectUser    `json:"project_user" gorm:"save_associations:false" valid:"-"`
}

// BeforeDelete - gorm hook, validate removal here & delete other related records (no cascade tag atm)
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	if err = tx.Where("owner_id = ? and owner_type = ?", c.ID, "comments").Delete(AttachedFile{}).Error; err != nil {
		return
	}
	return
}

// BeforeUpdate - gorm hook, fired before record update
func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	//check if original user_id is not being changed
	com := Comment{}
	if tx.Where("id = ? and user_id = ?", c.ID, c.UserID).First(&com); com.ID == 0 {
		return errors.New(helpers.NotFoundOrOwned("Comment"))
	}
	//delete removed file attachments
	ids := make([]uint64, len(c.AttachedFiles)+1) //force atleast 1 element for query to work... :/
	for i := 0; i < len(c.AttachedFiles); i++ {
		ids[i] = c.AttachedFiles[i].ID
	}
	if err = tx.Where("owner_id = ? and owner_type = ? and id not in(?)", c.ID, "comments", ids).Delete(AttachedFile{}).Error; err != nil {
		return
	}
	return
}

// BeforeSave - gorm hook, fired before record create or update
func (c *Comment) BeforeSave(tx *gorm.DB) (err error) {
	//check if task user is the same as comment user
	task := Task{}
	if tx.Where("id = ? and user_id = ?", c.TaskID, c.UserID).First(&task); task.ID == 0 {
		return errors.New(helpers.NotFoundOrOwned("Task"))
	}
	if c.IsSolution {
		//mark task as completed
		if err = tx.Model(Task{}).Where("id = ? and user_id = ?", c.TaskID, c.UserID).UpdateColumn("completed", true).Error; err != nil {
			return
		}
	}
	return
}
