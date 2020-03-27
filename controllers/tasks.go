package controllers

import (
	"net/http"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//tasksGet handles get all tasks request
func tasksGet(c *gin.Context) {
	var tasks []models.Task
	query := models.DB.Where("user_id = ?", currentUserID(c)).
		Preload("Project").Preload("Category")
	query = query.Preload("TaskLogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("session_id = 0 and minutes > 0")
	})
	query.Order("tasks.completed asc, tasks.created_at asc").Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

//taskGet handles get task request
func taskGet(c *gin.Context) {
	id := c.Param("id")
	task := models.Task{}
	query := models.DB.Where("user_id = ?", currentUserID(c)).
		Preload("AttachedFiles").Preload("Category")
	query = query.Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at asc")
	})
	query = query.Preload("TaskLogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("session_id = 0 and minutes > 0")
	})
	query.Preload("Comments.AttachedFiles").First(&task, id)
	if task.ID == 0 {
		abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Task"))
		return
	}
	c.JSON(http.StatusOK, task)
}

//taskNewGet handles get new task request
func taskNewGet(c *gin.Context) {
	projectID := helpers.AtoUint64(c.Query("project_id"))
	vm := models.EditTaskVM{}
	userID := currentUserID(c)
	models.DB.Where("user_id = ?", userID).Find(&vm.Projects)
	models.DB.Where("user_id = ?", userID).Find(&vm.Categories)
	vm.Task.Priority = models.PRIORITY4
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
	vm.Task.StartDate = time.Now()
	c.JSON(http.StatusOK, vm)
}

//taskEditGet handles get edit task request
func taskEditGet(c *gin.Context) {
	vm := models.EditTaskVM{}
	userID := currentUserID(c)
	models.DB.Where("user_id = ?", userID).Find(&vm.Projects)
	models.DB.Where("user_id = ?", userID).Preload("AttachedFiles").Preload("Project").Preload("Periodicity").First(&vm.Task, c.Param("id"))
	models.DB.Where("user_id = ?", userID).Find(&vm.Categories)
	c.JSON(http.StatusOK, vm)
}

//tasksPost handles create task request
func tasksPost(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if task.Priority == 0 {
		task.Priority = models.PRIORITY4
	}
	task.Completed = false
	task.UserID = currentUserID(c)
	task.Periodicity.UserID = currentUserID(c)
	if err := models.DB.Create(&task).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksPut handles update task request
func tasksPut(c *gin.Context) {
	//id := c.Param("id")
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	task.UserID = currentUserID(c)
	task.Periodicity.UserID = currentUserID(c)
	if err := models.DB.Save(&task).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksDelete handles delete task request
func tasksDelete(c *gin.Context) {
	id := c.Param("id")
	task := models.Task{}
	models.DB.Where("user_id = ?", currentUserID(c)).First(&task, id)
	if task.ID == 0 {
		abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Task"))
		return
	}
	if err := models.DB.Delete(&task).Error; err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//tasksSummaryGet handles get tasks statistics request
func tasksSummaryGet(c *gin.Context) {
	vm := models.TasksSummaryVM{}
	userID := currentUserID(c)
	if err := models.DB.Model(models.Task{}).Where("user_id = ?", userID).Count(&vm.Count).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	if err := models.DB.Where("user_id = ?", userID).Order("id desc").Limit(5).Find(&vm.LatestTasks).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	if err := models.DB.Where("user_id = ? and minutes > 0", userID).Order("id desc").Limit(5).Preload("Task").Find(&vm.LatestTaskLogs).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}
