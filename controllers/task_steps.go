package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//taskStepsGetHandler handles get all task steps request
func taskStepsGetHandler(c *gin.Context) {
	var steps []models.TaskStep
	models.DB.Order("order").Find(&steps)
	c.JSON(http.StatusOK, steps)
}

//taskStepGetHandler handles get task step request
func taskStepGetHandler(c *gin.Context) {
	id := c.Param("id")
	step := models.TaskStep{}
	models.DB.First(&step, id)
	if step.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task step not found"})
		return
	}
	c.JSON(http.StatusOK, step)
}

//taskStepsPostHandler handles create task step request
func taskStepsPostHandler(c *gin.Context) {
	step := models.TaskStep{}
	if err := c.ShouldBindJSON(&step); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&step).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//taskStepsPutHandler handles update task step request
func taskStepsPutHandler(c *gin.Context) {
	//id := c.Param("id")
	step := models.TaskStep{}
	if err := c.ShouldBindJSON(&step); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&step).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//taskStepsDeleteHandler handles delete task step request
func taskStepsDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	step := models.TaskStep{}
	models.DB.First(&step, id)
	if step.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task step not found"})
		return
	}
	if err := models.DB.Delete(&step).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
