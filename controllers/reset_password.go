package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/denisbakhtin/projectmanager/services"
	"github.com/gin-gonic/gin"
)

//forgotPost handles password reset request
func forgotPost(c *gin.Context) {
	vm := models.ForgotVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	models.DB.Where("email = ?", strings.ToLower(vm.Email)).First(&user)
	if user.ID == 0 {
		abortWithError(c, http.StatusBadRequest, fmt.Errorf("User not found"))
		return
	}

	user.Token = helpers.CreateSecureToken()
	if err := models.DB.Save(&user).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	go services.SendPasswordResetMessage(c, &user)

	c.JSON(http.StatusOK, gin.H{})
}

//resetPost handles password reset request
func resetPost(c *gin.Context) {
	vm := models.ResetVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	models.DB.Where("token = ?", vm.Token).First(&user)
	if user.ID == 0 {
		abortWithError(c, http.StatusBadRequest, fmt.Errorf("User not found"))
		return
	}
	user.Token = ""
	user.PasswordHash = helpers.CreatePasswordHash(vm.Password)
	if err := models.DB.Save(&user).Error; err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := user.CreateJWTToken(); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	go services.SendPasswordResetConfirmation(c, &user)

	c.JSON(200, gin.H{"token": user.Token})
}
