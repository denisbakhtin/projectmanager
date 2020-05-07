package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//TasksDB is a tasks db repository
var TasksDB TasksRepository

func init() {
	TasksDB = &tasksRepository{}
}

//TasksRepository is a repository of tasks
type TasksRepository interface {
	GetAll(userID uint64) ([]Task, error)
	Get(userID uint64, id interface{}) (Task, error)
	GetNew(userID uint64, projectID uint64) (EditTaskVM, error)
	GetEdit(userID uint64, id interface{}) (EditTaskVM, error)
	Create(userID uint64, task Task) (Task, error)
	Update(userID uint64, task Task) (Task, error)
	Delete(userID uint64, id interface{}) error
	Summary(userID uint64) (TasksSummaryVM, error)
}

type tasksRepository struct{}

//GetAll returns all tasks owned by specified user
func (tr *tasksRepository) GetAll(userID uint64) ([]Task, error) {
	var tasks []Task
	query := db.Where("user_id = ?", userID).Preload("Project").Preload("Category")
	query = query.Preload("TaskLogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("session_id = 0 and minutes > 0")
	})
	err := query.Order("tasks.completed asc, COALESCE(end_date, CURRENT_DATE) < CURRENT_DATE desc, priority asc").Find(&tasks).Error
	return tasks, err
}

//Get fetches a task by its id
func (tr *tasksRepository) Get(userID uint64, id interface{}) (Task, error) {
	task := Task{}
	query := db.Where("user_id = ?", userID).Preload("AttachedFiles").Preload("Category")
	query = query.Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at asc")
	})
	query = query.Preload("TaskLogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("session_id = 0 and minutes > 0")
	})
	err := query.Preload("Comments.AttachedFiles").First(&task, id).Error
	return task, err
}

//GetNew returns a view model for creating a new task
func (tr *tasksRepository) GetNew(userID uint64, projectID uint64) (EditTaskVM, error) {
	vm := EditTaskVM{}
	if err := db.Where("user_id = ?", userID).Find(&vm.Projects).Error; err != nil {
		return EditTaskVM{}, err
	}
	if err := db.Where("user_id = ?", userID).Find(&vm.Categories).Error; err != nil {
		return EditTaskVM{}, err
	}
	vm.Task.Priority = PRIORITY4
	vm.Task.ProjectID = projectID
	if projectID != 0 {
		//set category_id same as in the project
		for i := range vm.Projects {
			if vm.Projects[i].ID == projectID {
				vm.Task.CategoryID = vm.Projects[i].CategoryID
			}
		}
	}
	if projectID == 0 && len(vm.Projects) > 0 {
		vm.Task.ProjectID = vm.Projects[0].ID
	}
	vm.Task.Periodicity.Weekdays = 0b1111111 //Mon == 1 .. Sun == 0000001
	now := time.Now()
	vm.Task.StartDate = &now

	return vm, nil
}

//GetEdit returns a view model for task edition
func (tr *tasksRepository) GetEdit(userID uint64, id interface{}) (EditTaskVM, error) {
	vm := EditTaskVM{}
	if err := db.Where("user_id = ?", userID).Find(&vm.Projects).Error; err != nil {
		return EditTaskVM{}, err
	}
	if err := db.Where("user_id = ?", userID).Preload("AttachedFiles").Preload("Project").Preload("Periodicity").First(&vm.Task, id).Error; err != nil {
		return EditTaskVM{}, err
	}
	if err := db.Where("user_id = ?", userID).Find(&vm.Categories).Error; err != nil {
		return EditTaskVM{}, err
	}

	return vm, nil
}

//Create treates a new category in db
func (tr *tasksRepository) Create(userID uint64, task Task) (Task, error) {
	task.UserID = userID
	if task.Priority == 0 {
		task.Priority = PRIORITY4
	}
	task.Completed = false
	task.Periodicity.UserID = userID
	err := db.Create(&task).Error
	return task, err
}

//Update updates a category in db
func (tr *tasksRepository) Update(userID uint64, task Task) (Task, error) {
	task.UserID = userID
	task.Periodicity.UserID = userID
	err := db.Save(&task).Error
	return task, err
}

//Delete removes a category from db
func (tr *tasksRepository) Delete(userID uint64, id interface{}) error {
	task := Task{}
	if err := db.Where("user_id = ?", userID).First(&task, id).Error; err != nil {
		return err
	}
	if err := db.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

//Summary returns summary info for a dashboard
func (tr *tasksRepository) Summary(userID uint64) (TasksSummaryVM, error) {
	vm := TasksSummaryVM{}
	if err := db.Model(Task{}).Where("user_id = ?", userID).Count(&vm.Count).Error; err != nil {
		return TasksSummaryVM{}, err
	}
	if err := db.Where("user_id = ?", userID).Order("id desc").Limit(5).Find(&vm.LatestTasks).Error; err != nil {
		return TasksSummaryVM{}, err
	}
	if err := db.Where("user_id = ? and minutes > 0", userID).Order("id desc").Limit(5).Preload("Task").Find(&vm.LatestTaskLogs).Error; err != nil {
		return TasksSummaryVM{}, err
	}
	return vm, nil
}
