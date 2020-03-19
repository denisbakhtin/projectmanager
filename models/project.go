package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/jinzhu/gorm"
)

//Project represents a record from projects table
type Project struct {
	ID            uint64         `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Favorite      bool           `json:"favorite"`
	Name          string         `json:"name" valid:"required,length(1|1500)"`
	Description   string         `json:"description" valid:"length(0|100000)"`
	Archived      bool           `json:"archived"`
	UserID        uint64         `json:"user_id" valid:"-"`
	User          User           `json:"user" gorm:"save_associations:false" valid:"-"`
	CategoryID    uint64         `json:"category_id,string,omitempty"`
	AttachedFiles []AttachedFile `json:"files" gorm:"polymorphic:Owner" valid:"-"`
	Tasks         []Task         `json:"tasks" gorm:"save_associations:false" valid:"-"`
	Category      Category       `json:"category" gorm:"save_associations:false" valid:"-"`
	//ProjectUsers  []ProjectUser  `json:"project_users" valid:"-"`
}

// BeforeDelete - gorm hook, validate removal here & delete other related records (no cascade tag atm)
func (p *Project) BeforeDelete(tx *gorm.DB) (err error) {
	//check existence of important associated records and prohibit if any
	if tx.Model(p).Association("Tasks").Count() > 0 {
		return fmt.Errorf("Project has associated tasks with it")
	}
	if err = tx.Where("owner_id = ? and owner_type = ?", p.ID, "projects").Delete(AttachedFile{}).Error; err != nil {
		return
	}
	/*
		if err = tx.Where("project_id = ?", p.ID).Delete(ProjectUser{}).Error; err != nil {
			return
		}
	*/
	return
}

// BeforeSave - gorm hook, fired before record update or creation
func (p *Project) BeforeSave(tx *gorm.DB) (err error) {
	//check category belongs to the same user ^_^
	if p.CategoryID > 0 {
		cat := Category{}
		if tx.Where("user_id = ?", p.UserID).First(&cat, p.CategoryID); cat.ID == 0 {
			return errors.New(helpers.NotFoundOrOwned("Category"))
		}
	}
	return
}

// BeforeUpdate - gorm hook, fired before record update
func (p *Project) BeforeUpdate(tx *gorm.DB) (err error) {
	//check if original user_id is not being changed
	proj := Project{}
	if tx.Where("id = ? and user_id = ?", p.ID, p.UserID).First(&proj); proj.ID == 0 {
		return errors.New(helpers.NotFoundOrOwned("Project"))
	}
	//delete removed file attachments
	ids := make([]uint64, len(p.AttachedFiles))
	for i := 0; i < len(p.AttachedFiles); i++ {
		ids[i] = p.AttachedFiles[i].ID
	}
	if err = tx.Where("owner_id = ? and owner_type = ? and id not in(?)", p.ID, "projects", ids).Delete(AttachedFile{}).Error; err != nil {
		return
	}
	/*
		userIds := make([]uint64, len(project.ProjectUsers)) //add atleast one non-existent id for query to work :)
		for i := 0; i < len(project.ProjectUsers); i++ {
			userIds[i] = project.ProjectUsers[i].ID
		}
		if err := models.DB.Where("project_id = ? and id not in (?)", project.ID, userIds).Delete(models.ProjectUser{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	*/
	return
}

//EditProjectVM is a view model for a new or an edited project
type EditProjectVM struct {
	Project    `json:"project"`
	Categories []Category `json:"categories"`
}

//ProjectsSummaryVM is a view model for projects statistics
type ProjectsSummaryVM struct {
	Count int `json:"count"`
}
