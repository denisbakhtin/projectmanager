package controllers

import (
	"fmt"
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
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
	user.Name = vm.Name
	if len(vm.CurrentPassword) > 0 && len(vm.NewPassword) > 0 {
		if !user.HasPassword(vm.CurrentPassword) {
			abortWithError(c, http.StatusBadRequest, fmt.Errorf("Wrong current password"))
			return
		}
		if err := helpers.CheckNewPassword(vm.NewPassword); err != nil {
			abortWithError(c, http.StatusBadRequest, err)
			return
		}
		user.PasswordHash = helpers.CreatePasswordHash(vm.NewPassword)
	}
	if err := models.DB.Save(&user).Error; err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := user.CreateJWTToken(); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{"token": user.Token})
}
