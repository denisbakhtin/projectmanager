package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//spentGet handles spent report get request
func spentGet(c *gin.Context) {
	var taskLogs []models.TaskLog
	models.DB.Preload("Task.Project").Preload("Task").
		Preload("Task.Category").
		Where("session_id = 0 and minutes > 0 and user_id = ?", currentUserID(c)).
		Find(&taskLogs)
	c.JSON(http.StatusOK, taskLogs)
}
