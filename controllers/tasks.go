package controllers

import (
	"net/http"

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
		c.JSON(http.StatusNotFound, helpers.NotFoundOrOwned("Task"))
		return
	}
	c.JSON(http.StatusOK, task)
}

//taskNewGet handles get new task request
func taskNewGet(c *gin.Context) {
	vm := models.EditTaskVM{}
	userID := currentUserID(c)
	models.DB.Where("user_id = ?", userID).Find(&vm.Projects)
	models.DB.Where("user_id = ?", userID).Find(&vm.Categories)
	vm.Task.Priority = models.PRIORITY4
	c.JSON(http.StatusOK, vm)
}

//taskEditGet handles get edit task request
func taskEditGet(c *gin.Context) {
	vm := models.EditTaskVM{}
	userID := currentUserID(c)
	models.DB.Where("user_id = ?", userID).Find(&vm.Projects)
	models.DB.Where("user_id = ?", userID).Preload("AttachedFiles").Preload("Project").First(&vm.Task, c.Param("id"))
	models.DB.Where("user_id = ?", userID).Find(&vm.Categories)
	c.JSON(http.StatusOK, vm)
}

//tasksPost handles create task request
func tasksPost(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if task.Priority == 0 {
		task.Priority = models.PRIORITY4
	}
	task.Completed = false
	task.UserID = currentUserID(c)
	if err := models.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksPut handles update task request
func tasksPut(c *gin.Context) {
	//id := c.Param("id")
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	task.UserID = currentUserID(c)
	if err := models.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusNotFound, helpers.NotFoundOrOwned("Task"))
		return
	}
	if err := models.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
