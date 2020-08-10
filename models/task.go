package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/jinzhu/gorm"
)

//Task priorities from high to low, their proper titles are set on client side
const (
	PRIORITY1 = iota + 1 //== 1
	PRIORITY2            // == 2, etc
	PRIORITY3
	PRIORITY4
)

//Task represents a row from tasks table
type Task struct {
	ID            uint64         `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"treated_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Name          string         `json:"name" valid:"required,length(1|1500)"`
	Destription   string         `json:"description" valid:"length(0|10000)"`
	StartDate     *time.Time     `json:"start_date"`
	EndDate       *time.Time     `json:"end_date"`
	PeriodicityID uint64         `json:"periodicity_id,omitempty" gorm:"index"`
	Periodicity   Periodicity    `json:"periodicity" gorm:"save_associations:true"`
	Completed     bool           `json:"completed"`
	Priority      uint           `json:"priority,string"` //see constants
	ProjectID     uint64         `json:"project_id,string,omitempty" gorm:"index" valid:"required"`
	Project       Project        `json:"project" gorm:"save_associations:false" valid:"-"`
	CategoryID    uint64         `json:"category_id,string,omitempty" gorm:"index"`
	Category      Category       `json:"category" gorm:"save_associations:false" valid:"-"`
	UserID        uint64         `json:"user_id" valid:"-"`
	User          User           `json:"user" gorm:"save_associations:false" valid:"-"`
	Comments      []Comment      `json:"comments" gorm:"save_associations:false" valid:"-"`
	TaskLogs      []TaskLog      `json:"task_logs" gorm:"save_associations:false" valid:"-"`
	AttachedFiles []AttachedFile `json:"files" gorm:"polymorphic:Owner" valid:"-"`
	//ProjectUserID uint64         `json:"project_user_id" gorm:"index" valid:"required"`
	//ProjectUser   ProjectUser    `json:"project_user" gorm:"save_associations:false" valid:"-"`
}

// BeforeDelete - gorm hook, validate removal here & delete other related records (no cascade tag atm)
func (t *Task) BeforeDelete(tx *gorm.DB) (err error) {
	//check existence of important associated records and prohibit if any
	if tx.Model(t).Association("TaskLogs").Count() > 0 {
		return fmt.Errorf("Task has associated activities with it")
	}
	if err = tx.Where("owner_id = ? and owner_type = ?", t.ID, "tasks").Delete(AttachedFile{}).Error; err != nil {
		return
	}
	if err = tx.Where("task_id = ?", t.ID).Delete(Comment{}).Error; err != nil {
		return
	}
	return
}

// BeforeSave - gorm hook, fired before record update or treation
func (t *Task) BeforeSave(tx *gorm.DB) (err error) {
	//check category & project belongs to the same user ^_^
	if t.CategoryID > 0 {
		cat := Category{}
		if tx.Where("user_id = ?", t.UserID).First(&cat, t.CategoryID); cat.ID == 0 {
			return errors.New(helpers.NotFoundOrOwned("Category"))
		}
	}
	proj := Project{}
	if tx.Where("user_id = ?", t.UserID).First(&proj, t.ProjectID); proj.ID == 0 {
		return errors.New(helpers.NotFoundOrOwned("Project"))
	}
	return
}

// BeforeUpdate - gorm hook, fired before record update
func (t *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	//check if original user_id is not being changed
	task := Task{}
	if tx.Where("id = ? and user_id = ?", t.ID, t.UserID).First(&task); task.ID == 0 {
		return errors.New(helpers.NotFoundOrOwned("Task"))
	}
	//delete removed file attachments
	ids := make([]uint64, len(t.AttachedFiles)+1) //force atleast 1 element for query to work... :/
	for i := 0; i < len(t.AttachedFiles); i++ {
		ids[i] = t.AttachedFiles[i].ID
	}
	if err = tx.Where("owner_id = ? and owner_type = ? and id not in(?)", t.ID, "tasks", ids).Delete(AttachedFile{}).Error; err != nil {
		return
	}
	return
}
