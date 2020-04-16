package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//usersGet handles get all users request
func usersGet(c *gin.Context) {
	users, err := models.UsersDB.GetAll()
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

//userGet handles get user request
func userGet(c *gin.Context) {
	user, err := models.UsersDB.Get(c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("User"))
		} else {
			abortWithError(c, http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

//userStatusPut handles user status update
func userStatusPut(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.UsersDB.UpdateStatus(user); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//usersSummaryGet handles get users statistics request
func usersSummaryGet(c *gin.Context) {
	vm, err := models.UsersDB.Summary()
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}
