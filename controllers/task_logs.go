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
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.TaskLogsDB.Create(currentUserID(c), taskLog); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//taskLogsPut handles update taskLog request
func taskLogsPut(c *gin.Context) {
	taskLog := models.TaskLog{}
	if err := c.ShouldBindJSON(&taskLog); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.TaskLogsDB.Update(currentUserID(c), taskLog); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
