package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/denisbakhtin/projectmanager/viewmodels"
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
	vm := viewmodels.AccountVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	if u, exists := c.Get("user"); exists {
		user = u.(models.User)
	}
	user.Name = vm.Name
	if len(vm.CurrentPassword) > 0 && len(vm.NewPassword) > 0 {
		if !user.HasPassword(vm.CurrentPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong current password"})
			return
		}
		if err := helpers.CheckNewPassword(vm.NewPassword); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.PasswordHash = helpers.CreatePasswordHash(vm.NewPassword)
	}
	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.CreateJWTToken(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": user.Token})
}
