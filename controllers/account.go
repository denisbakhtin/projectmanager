package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//accountGet handles user account request
func accountGet(c *gin.Context) {
	user := models.User{}
	if u, exists := c.Get("user"); exists {
		user = u.(models.User)
	}

	c.JSON(200, user)
}

//accountPut handles user account update
func accountPut(c *gin.Context) {
	vm := models.AccountVM{}

	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	if u, exists := c.Get("user"); exists {
		user = u.(models.User)
	}
	user, err := models.UsersDB.UpdateAccount(vm, user)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200, gin.H{"token": user.JWTToken})
}
