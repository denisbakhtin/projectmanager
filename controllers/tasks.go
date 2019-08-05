package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//tasksGetHandler handles get all tasks request
func tasksGetHandler(c *gin.Context) {
	var tasks []models.Task
	models.DB.Preload("ProjectUser").Preload("Project").Preload("TaskStep").Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

//taskGetHandler handles get task request
func taskGetHandler(c *gin.Context) {
	id := c.Param("id")
	task := models.Task{}
	models.DB.First(&task, id)
	if task.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

//tasksPostHandler handles create role request
func tasksPostHandler(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}
	if u, exists := c.Get("user"); exists {
		user = u.(models.User)
	}
	puser := models.ProjectUser{}
	models.DB.Where("user_id = ?", user.ID).First(&puser)
	if puser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project user not found"})
		return
	}
	task.ProjectUserID = puser.ID
	step := models.TaskStep{}
	models.DB.Order("order").First(&step)
	task.TaskStepID = step.ID
	if err := models.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksPutHandler handles update task request
func tasksPutHandler(c *gin.Context) {
	//id := c.Param("id")
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksDeleteHandler handles delete role request
func tasksDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	task := models.Task{}
	models.DB.First(&task, id)
	if task.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	if err := models.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
