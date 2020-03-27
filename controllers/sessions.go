package controllers

import (
	"fmt"
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//sessionsGet handles get all sessions request
func sessionsGet(c *gin.Context) {
	var sessions []models.Session
	models.DB.Where("user_id = ?", currentUserID(c)).Preload("TaskLogs").Find(&sessions)
	c.JSON(http.StatusOK, sessions)
}

//sessionGet handles get session request
func sessionGet(c *gin.Context) {
	id := c.Param("id")
	session := models.Session{}
	userID := currentUserID(c)
	models.DB.Where("user_id = ?", userID).
		Preload("TaskLogs").Preload("TaskLogs.Task").Preload("TaskLogs.Task.Project").
		First(&session, id)
	if session.ID == 0 {
		abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Session"))
		return
	}
	c.JSON(http.StatusOK, session)
}

//sessionNewGet handles get new session request
func sessionNewGet(c *gin.Context) {
	var logs []models.TaskLog
	userID := currentUserID(c)
	models.DB.
		Where("user_id = ? and minutes > 0 and session_id = 0", userID).
		Preload("Task").Preload("Task.Project").Find(&logs)
	c.JSON(http.StatusOK, logs)
}

//sessionsPost handles create session request
func sessionsPost(c *gin.Context) {
	session := models.Session{}
	if err := c.ShouldBindJSON(&session); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	userID := currentUserID(c)
	session.UserID = userID

	if err := models.DB.Create(&session).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//sessionsDelete handles delete session request
func sessionsDelete(c *gin.Context) {
	id := c.Param("id")
	session := models.Session{}
	models.DB.Preload("TaskLogs").Where("user_id = ?", currentUserID(c)).First(&session, id)
	if session.ID == 0 {
		abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Session"))
		return
	}
	if len(session.TaskLogs) > 0 {
		abortWithError(c, http.StatusBadRequest, fmt.Errorf("Can not remove non-empty session"))
		return
	}
	if err := models.DB.Delete(&session).Error; err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//sessionsSummaryGet handles get sessions statistics request
func sessionsSummaryGet(c *gin.Context) {
	vm := models.SessionsSummaryVM{}
	userID := currentUserID(c)
	if err := models.DB.Model(models.Session{}).Where("user_id = ?", userID).Count(&vm.Count).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}
