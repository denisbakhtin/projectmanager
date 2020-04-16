package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//spentGet handles spent report get request
func spentGet(c *gin.Context) {
	taskLogs, err := models.ReportsDB.Spent(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, taskLogs)
}
