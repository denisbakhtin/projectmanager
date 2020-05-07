package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestProjectsRepositoryGetAll(t *testing.T) {
	getOrCreateUnrelatedProject()
	list, err := ProjectsDB.GetAll(userID)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 1)
	for _, com := range list {
		assert.Equal(t, com.UserID, userID, "Project should belong to the specified user")
	}
}

func TestProjectsRepositoryGet(t *testing.T) {
	project := getOrCreateUnrelatedProject()
	p, err := ProjectsDB.Get(userID, project.ID)
	assert.Nil(t, err)
	assert.NotZero(t, p.ID)
	assert.Equal(t, p.UserID, userID)
}

func TestProjectsRepositoryCreate(t *testing.T) {
	name := fmt.Sprintf("Project-%d", time.Now().Nanosecond())
	project := Project{
		Name: name,
	}

	all, _ := ProjectsDB.GetAll(userID)
	count := len(all)
	p, err := ProjectsDB.Create(userID, project)
	assert.Nil(t, err)
	assert.NotZero(t, p.ID)
	assert.Equal(t, p.Name, name)

	p, err = ProjectsDB.Get(userID, p.ID)
	assert.Nil(t, err)
	assert.NotZero(t, p.ID)
	assert.Equal(t, p.Name, name)

	all, _ = ProjectsDB.GetAll(userID)
	assert.Equal(t, count+1, len(all), "The number of projects should have been increased by 1")
}

func TestProjectsRepositoryUpdate(t *testing.T) {
	name := fmt.Sprintf("Project-%d", time.Now().Nanosecond())
	project := Project{}
	db.Where("user_id = ?", userID).First(&project)
	assert.NotZero(t, project.ID)
	assert.NotEqual(t, project.Name, name)

	project.Name = name
	p, err := ProjectsDB.Update(userID, project)
	assert.Nil(t, err)
	assert.Equal(t, p.UserID, userID)
	assert.Equal(t, p.Name, name)
}

func TestProjectsRepositoryDelete(t *testing.T) {
	err := ProjectsDB.Delete(userID, 11111111111)
	assert.NotNil(t, err)
	assert.True(t, gorm.IsRecordNotFoundError(err), "This project should not exist")

	project := getOrCreateUnrelatedProject()
	assert.NotZero(t, project.ID)

	err = ProjectsDB.Delete(userID+1, project.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Should check project owner")

	err = ProjectsDB.Delete(userID, project.ID)
	assert.Nil(t, err)
	_, err = ProjectsDB.Get(userID, project.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Project should have been removed")
}

func TestProjectsRepositoryToggleFavorite(t *testing.T) {
	p := getOrCreateUnrelatedProject()
	p.Favorite = !p.Favorite
	pr, err := ProjectsDB.ToggleFavorite(userID, p)
	assert.Nil(t, err)
	proj, err := ProjectsDB.Get(userID, p.ID)
	assert.Nil(t, err)
	assert.NotZero(t, proj.ID)
	assert.Equal(t, pr.ID, proj.ID)
	assert.Equal(t, pr.Favorite, proj.Favorite)
	assert.Equal(t, p.Name, proj.Name)
	assert.Equal(t, p.Favorite, proj.Favorite)
}

func TestProjectsRepositoryToggleArchive(t *testing.T) {
	p := getOrCreateUnrelatedProject()
	p.Archived = !p.Archived
	pr, err := ProjectsDB.ToggleArchive(userID, p)
	assert.Nil(t, err)
	proj, err := ProjectsDB.Get(userID, p.ID)
	assert.Nil(t, err)
	assert.NotZero(t, proj.ID)
	assert.Equal(t, pr.ID, proj.ID)
	assert.Equal(t, pr.Archived, proj.Archived)
	assert.Equal(t, p.Name, proj.Name)
	assert.Equal(t, p.Archived, proj.Archived)
}

func TestProjectsRepositoryGetAllFavorite(t *testing.T) {
	p := getOrCreateUnrelatedProject()
	p.Favorite = true
	_, err := ProjectsDB.ToggleFavorite(userID, p)
	assert.Nil(t, err)
	list, err := ProjectsDB.GetAllFavorite(userID)
	assert.Nil(t, err)
	assert.Greater(t, len(list), 0)
}

func TestProjectsRepositorySummary(t *testing.T) {
	//atleast one project exists
	getOrCreateUnrelatedProject()
	vm, err := ProjectsDB.Summary(userID)
	assert.Nil(t, err)
	assert.NotEmpty(t, vm)
	assert.Greater(t, vm.Count, 0)
}
