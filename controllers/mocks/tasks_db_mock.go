package mocks

import (
	"fmt"
	"strconv"
	"time"

	"github.com/denisbakhtin/projectmanager/models"
)

//TasksDBMock is a TasksDB repository mock
type TasksDBMock struct {
	Tasks []models.Task
}

func (tm *TasksDBMock) GetAll(userID uint64) ([]models.Task, error) {
	return tm.Tasks, nil
}

func (tm *TasksDBMock) Get(userID uint64, id interface{}) (models.Task, error) {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for _, c := range tm.Tasks {
		if c.ID == idi {
			return c, nil
		}
	}
	return models.Task{}, fmt.Errorf("Task not found")
}

//GetNew returns a view model for creating a new task
func (tm *TasksDBMock) GetNew(userID uint64, projectID uint64) (models.EditTaskVM, error) {
	vm := models.EditTaskVM{}
	vm.Task.Priority = models.PRIORITY4
	vm.Task.ProjectID = 1
	now := time.Now()
	vm.Task.StartDate = &now
	return vm, nil
}

func (tm *TasksDBMock) GetEdit(userID uint64, id interface{}) (models.EditTaskVM, error) {
	var err error
	vm := models.EditTaskVM{}
	vm.Task, err = models.TasksDB.Get(userID, "1")
	if err != nil {
		return models.EditTaskVM{}, err
	}

	return vm, nil
}

//Create tmeates a new task in db
func (tm *TasksDBMock) Create(userID uint64, task models.Task) (models.Task, error) {
	task.UserID = userID
	tm.Tasks = append(tm.Tasks, task)
	return task, nil
}

//Update updates a task in db
func (tm *TasksDBMock) Update(userID uint64, task models.Task) (models.Task, error) {
	task.UserID = userID
	for i := range tm.Tasks {
		if tm.Tasks[i].ID == task.ID {
			tm.Tasks[i] = task
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("Task not found")
}

//Delete removes a task from db
func (tm *TasksDBMock) Delete(userID uint64, id interface{}) error {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range tm.Tasks {
		if tm.Tasks[i].ID == idi && tm.Tasks[i].UserID == userID {
			tm.Tasks[i] = tm.Tasks[len(tm.Tasks)-1] // Copy last element to index i.
			tm.Tasks = tm.Tasks[:len(tm.Tasks)-1]   // Truncate slice.
			return nil
		}
	}
	return nil
}

func (tm *TasksDBMock) Summary(userID uint64) (models.TasksSummaryVM, error) {
	return models.TasksSummaryVM{Count: len(tm.Tasks)}, nil
}
