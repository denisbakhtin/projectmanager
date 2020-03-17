package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//taskLogsPost handles create taskLog request
func taskLogsPost(c *gin.Context) {
	taskLog := models.TaskLog{}
	if err := c.ShouldBindJSON(&taskLog); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	taskLog.UserID = currentUserID(c)
	taskLog.SessionID = 0
	if err := models.DB.Create(&taskLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, taskLog)
}

//taskLogsPut handles update taskLog request
func taskLogsPut(c *gin.Context) {
	//id := c.Param("id")
	taskLog := models.TaskLog{}
	if err := c.ShouldBindJSON(&taskLog); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	taskLog.UserID = currentUserID(c)
	taskLog.SessionID = 0
	if err := models.DB.Save(&taskLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, taskLog)
}
