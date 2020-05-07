package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestTasksRepositoryGetAll(t *testing.T) {
	getOrCreateTask()
	list, err := TasksDB.GetAll(userID)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 1)
	for _, com := range list {
		assert.Equal(t, com.UserID, userID, "Task should belong to the specified user")
	}
}

func TestTasksRepositoryGet(t *testing.T) {
	task := getOrCreateTask()
	p, err := TasksDB.Get(userID, task.ID)
	assert.Nil(t, err)
	assert.NotZero(t, p.ID)
	assert.Equal(t, p.UserID, userID)
}

func TestTasksRepositoryGetNew(t *testing.T) {
	project := getOrCreateUnrelatedProject()
	vm, err := TasksDB.GetNew(userID, project.ID)
	assert.Nil(t, err)
	assert.NotZero(t, len(vm.Projects))
	assert.Zero(t, vm.Task.ID)
	assert.Equal(t, vm.Task.CategoryID, project.CategoryID)
	assert.Equal(t, vm.Task.ProjectID, project.ID)
}

func TestTasksRepositoryGetEdit(t *testing.T) {
	task := getOrCreateTask()
	getOrCreateUnrelatedProject()
	vm, err := TasksDB.GetEdit(userID, task.ID)
	assert.Nil(t, err)
	assert.NotZero(t, len(vm.Projects))
	assert.NotZero(t, vm.Task.ID)
	assert.Equal(t, vm.Task.ID, task.ID)
}

func TestTasksRepositoryCreate(t *testing.T) {
	project := getOrCreateUnrelatedProject()
	name := fmt.Sprintf("Task-%d", time.Now().Nanosecond())
	task := Task{
		Name:      name,
		ProjectID: project.ID,
	}

	all, _ := TasksDB.GetAll(userID)
	count := len(all)
	t1, err := TasksDB.Create(userID, task)
	assert.Nil(t, err)
	assert.NotZero(t, t1.ID)
	assert.Equal(t, t1.Name, name)

	t1, err = TasksDB.Get(userID, t1.ID)
	assert.Nil(t, err)
	assert.NotZero(t, t1.ID)
	assert.Equal(t, t1.Name, name)

	all, _ = TasksDB.GetAll(userID)
	assert.Equal(t, count+1, len(all), "The number of tasks should have been increased by 1")
}

func TestTasksRepositoryUpdate(t *testing.T) {
	name := fmt.Sprintf("Task-%d", time.Now().Nanosecond())
	task := Task{}
	db.Where("user_id = ?", userID).First(&task)
	assert.NotZero(t, task.ID)
	assert.NotEqual(t, task.Name, name)

	task.Name = name
	p, err := TasksDB.Update(userID, task)
	assert.Nil(t, err)
	assert.Equal(t, p.UserID, userID)
	assert.Equal(t, p.Name, name)
}

func TestTasksRepositoryDelete(t *testing.T) {
	err := TasksDB.Delete(userID, 11111111111)
	assert.NotNil(t, err)
	assert.True(t, gorm.IsRecordNotFoundError(err), "This task should not exist")

	task := Task{}
	err = db.Where("user_id = ? and NOT EXISTS(select null from task_logs where task_logs.task_id = tasks.id)", userID).First(&task).Error
	assert.Nil(t, err)
	assert.NotZero(t, task.ID)

	err = TasksDB.Delete(userID+1, task.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Should check task owner")

	err = TasksDB.Delete(userID, task.ID)
	assert.Nil(t, err)
	_, err = TasksDB.Get(userID, task.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Task should have been removed")
}

func TestTasksRepositorySummary(t *testing.T) {
	//atleast one task exists
	getOrCreateTask()
	vm, err := TasksDB.Summary(userID)
	assert.Nil(t, err)
	assert.NotEmpty(t, vm)
	assert.Greater(t, vm.Count, 0)
}
