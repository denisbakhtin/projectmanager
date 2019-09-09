package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//taskStepsGet handles get all task steps request
func taskStepsGet(c *gin.Context) {
	var steps []models.TaskStep
	models.DB.Order("order").Find(&steps)
	c.JSON(http.StatusOK, steps)
}

//taskStepGet handles get task step request
func taskStepGet(c *gin.Context) {
	id := c.Param("id")
	step := models.TaskStep{}
	models.DB.First(&step, id)
	if step.ID == 0 {
		c.JSON(http.StatusNotFound, "Task step not found")
		return
	}
	c.JSON(http.StatusOK, step)
}

//taskStepsPost handles create task step request
func taskStepsPost(c *gin.Context) {
	step := models.TaskStep{}
	if err := c.ShouldBindJSON(&step); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Create(&step).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//taskStepsPut handles update task step request
func taskStepsPut(c *gin.Context) {
	//id := c.Param("id")
	step := models.TaskStep{}
	if err := c.ShouldBindJSON(&step); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Save(&step).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//taskStepsDelete handles delete task step request
func taskStepsDelete(c *gin.Context) {
	id := c.Param("id")
	step := models.TaskStep{}
	models.DB.First(&step, id)
	if step.ID == 0 {
		c.JSON(http.StatusNotFound, "Task step not found")
		return
	}
	if err := models.DB.Delete(&step).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
