package models

import "github.com/jinzhu/gorm"

//ProjectsDB is a projects db repository
var ProjectsDB ProjectsRepository

func init() {
	ProjectsDB = &projectsRepository{}
}

//ProjectsRepository is a projects repository
type ProjectsRepository interface {
	GetAll(userID uint64) ([]Project, error)
	GetAllFavorite(userID uint64) ([]Project, error)
	Get(userID uint64, id interface{}) (Project, error)
	GetNew(userID uint64) (EditProjectVM, error)
	GetEdit(userID uint64, id interface{}) (EditProjectVM, error)
	Create(userID uint64, project Project) (Project, error)
	Update(userID uint64, project Project) (Project, error)
	ToggleArchive(userID uint64, project Project) (Project, error)
	ToggleFavorite(userID uint64, project Project) (Project, error)
	Delete(userID uint64, id interface{}) error
	Summary(userID uint64) (ProjectsSummaryVM, error)
}

type projectsRepository struct{}

//GetAll returns all projects owned by the specified user
func (pr *projectsRepository) GetAll(userID uint64) ([]Project, error) {
	var projects []Project
	err := DB.Preload("Tasks").Preload("Category").
		Where("user_id = ?", userID).Order("archived asc, created_at asc").
		Find(&projects).Error
	return projects, err
}

//GetAllFavorite returns all favorite projects owned by the specified user
func (pr *projectsRepository) GetAllFavorite(userID uint64) ([]Project, error) {
	var projects []Project
	err := DB.Select("id, name").Where("user_id = ? and favorite = true", userID).Order("created_at asc").Find(&projects).Error
	return projects, err
}

//Get fetches a project by its id
func (pr *projectsRepository) Get(userID uint64, id interface{}) (Project, error) {
	project := Project{}
	query := DB.Where("user_id = ?", userID).Preload("AttachedFiles").Preload("Category")
	query = query.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Order("tasks.completed asc, created_at asc")
	})
	query = query.Preload("Tasks.Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at asc")
	})
	query = query.Preload("Tasks.TaskLogs")

	err := query.First(&project, id).Error
	return project, err
}

//GetNew returns a viewmodel with data required for building a new project
func (pr *projectsRepository) GetNew(userID uint64) (EditProjectVM, error) {
	vm := EditProjectVM{}
	err := DB.Where("user_id = ?", userID).Find(&vm.Categories).Error
	return vm, err
}

//GetEdit returns a viewmodel with data required for editing the project
func (pr *projectsRepository) GetEdit(userID uint64, id interface{}) (EditProjectVM, error) {
	vm := EditProjectVM{}
	err := DB.Where("user_id = ?", userID).Find(&vm.Categories).Error
	if err != nil {
		return vm, err
	}
	err = DB.Preload("AttachedFiles").Preload("Tasks").First(&vm.Project, id).Error
	return vm, err
}

//Create preates a new project in db
func (pr *projectsRepository) Create(userID uint64, project Project) (Project, error) {
	project.UserID = userID
	err := DB.Create(&project).Error
	return project, err
}

//Update updates a project in db
func (pr *projectsRepository) Update(userID uint64, project Project) (Project, error) {
	project.UserID = userID
	err := DB.Save(&project).Error
	return project, err
}

//ToggleArchive toggles project's archived field
func (pr *projectsRepository) ToggleArchive(userID uint64, project Project) (Project, error) {
	err := DB.Model(&project).Where("user_id = ?", userID).UpdateColumn("archived", project.Archived).Error
	return project, err
}

//ToggleFavorite toggles project's favorite field
func (pr *projectsRepository) ToggleFavorite(userID uint64, project Project) (Project, error) {
	err := DB.Model(&project).Where("user_id = ?", userID).UpdateColumn("favorite", project.Favorite).Error
	return project, err
}

//Delete removes a project from db
func (pr *projectsRepository) Delete(userID uint64, id interface{}) error {
	project := Project{}
	if err := DB.Where("user_id = ?", userID).First(&project, id).Error; err != nil {
		return err
	}
	if err := DB.Delete(&project).Error; err != nil {
		return err
	}
	return nil
}

//Summary gets projects summary info for a dashboard
func (pr *projectsRepository) Summary(userID uint64) (ProjectsSummaryVM, error) {
	vm := ProjectsSummaryVM{}
	err := DB.Model(Project{}).Where("user_id = ?", userID).Count(&vm.Count).Error
	return vm, err
}
