package mocks

import (
	"fmt"
	"strconv"

	"github.com/denisbakhtin/projectmanager/models"
)

type ProjectsDBMock struct {
	Projects []models.Project
}

//GetAll returns all projects owned by the specified user
func (pr *ProjectsDBMock) GetAll(userID uint64) ([]models.Project, error) {
	return pr.Projects, nil
}

//GetAllFavorite returns all favorite projects owned by the specified user
func (pr *ProjectsDBMock) GetAllFavorite(userID uint64) ([]models.Project, error) {
	projects := make([]models.Project, 0, len(pr.Projects))
	for i := range pr.Projects {
		if pr.Projects[i].Favorite {
			projects = append(projects, pr.Projects[i])
		}
	}
	return projects, nil
}

//Get fetches a project by its id
func (pr *ProjectsDBMock) Get(userID uint64, id interface{}) (models.Project, error) {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range pr.Projects {
		if pr.Projects[i].ID == idi {
			return pr.Projects[i], nil
		}
	}
	return models.Project{}, fmt.Errorf("Project not found")
}

//GetNew returns a viewmodel with data required for building a new project
func (pr *ProjectsDBMock) GetNew(userID uint64) (models.EditProjectVM, error) {
	return models.EditProjectVM{}, nil
}

//GetEdit returns a viewmodel with data required for editing the project
func (pr *ProjectsDBMock) GetEdit(userID uint64, id interface{}) (models.EditProjectVM, error) {
	return models.EditProjectVM{}, nil
}

//Create preates a new project in db
func (pr *ProjectsDBMock) Create(userID uint64, project models.Project) (models.Project, error) {
	project.UserID = userID
	pr.Projects = append(pr.Projects, project)
	return project, nil
}

//Update updates a project in db
func (pr *ProjectsDBMock) Update(userID uint64, project models.Project) (models.Project, error) {
	for i := range pr.Projects {
		if pr.Projects[i].ID == project.ID {
			project.UserID = userID
			pr.Projects[i] = project
			return project, nil
		}
	}
	return models.Project{}, fmt.Errorf("Project not found")
}

//ToggleArchive toggles project's archived field
func (pr *ProjectsDBMock) ToggleArchive(userID uint64, project models.Project) (models.Project, error) {
	for i := range pr.Projects {
		if pr.Projects[i].ID == project.ID {
			pr.Projects[i].Archived = project.Archived
			return project, nil
		}
	}
	return models.Project{}, fmt.Errorf("Project not found")
}

//ToggleFavorite toggles project's favorite field
func (pr *ProjectsDBMock) ToggleFavorite(userID uint64, project models.Project) (models.Project, error) {
	for i := range pr.Projects {
		if pr.Projects[i].ID == project.ID {
			pr.Projects[i].Favorite = project.Favorite
			return project, nil
		}
	}
	return models.Project{}, fmt.Errorf("Project not found")
}

//Delete removes a project from db
func (pr *ProjectsDBMock) Delete(userID uint64, id interface{}) error {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range pr.Projects {
		if pr.Projects[i].ID == idi {
			pr.Projects[i] = pr.Projects[len(pr.Projects)-1] // Copy last element to index i.
			pr.Projects = pr.Projects[:len(pr.Projects)-1]   // Truncate slice.
			return nil
		}
	}
	return fmt.Errorf("Project not found")
}

//Summary gets projects summary info for a dashboard
func (pr *ProjectsDBMock) Summary(userID uint64) (models.ProjectsSummaryVM, error) {
	vm := models.ProjectsSummaryVM{Count: len(pr.Projects)}
	return vm, nil
}
