package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//sessionsGet handles get all sessions request
func sessionsGet(c *gin.Context) {
	sessions, err := models.SessionsDB.GetAll(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, sessions)
}

//sessionGet handles get session request
func sessionGet(c *gin.Context) {
	session, err := models.SessionsDB.Get(currentUserID(c), c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Page"))
		} else {
			abortWithError(c, http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, session)
}

//sessionNewGet handles get new session request
func sessionNewGet(c *gin.Context) {
	logs, err := models.SessionsDB.NewGet(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, logs)
}

//sessionsPost handles create session request
func sessionsPost(c *gin.Context) {
	session := models.Session{}
	if err := c.ShouldBindJSON(&session); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.SessionsDB.Create(currentUserID(c), session); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//sessionsDelete handles delete session request
func sessionsDelete(c *gin.Context) {
	if err := models.SessionsDB.Delete(currentUserID(c), c.Param("id")); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//sessionsSummaryGet handles get sessions statistics request
func sessionsSummaryGet(c *gin.Context) {
	vm, err := models.SessionsDB.Summary(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}
