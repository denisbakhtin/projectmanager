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
	taskLog, err := models.TaskLogsDB.Create(currentUserID(c), taskLog)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, taskLog)
}

//taskLogsPut handles update taskLog request
func taskLogsPut(c *gin.Context) {
	taskLog := models.TaskLog{}
	if err := c.ShouldBindJSON(&taskLog); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	taskLog, err := models.TaskLogsDB.Update(currentUserID(c), taskLog)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, taskLog)
}

//taskLogsLatestGet handles get latest task logs request
func taskLogsLatestGet(c *gin.Context) {
	logs, err := models.TaskLogsDB.Latest(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, logs)
}
