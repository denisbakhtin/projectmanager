package controllers

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//loginPost handles user login
func loginPost(c *gin.Context) {
	vm := models.LoginVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	user, err := models.UsersDB.Login(vm)
	if err != nil {
		abortWithError(c, http.StatusUnauthorized, err)
		return
	}
	c.JSON(200, gin.H{"token": user.JWTToken})
}

//activatePost handles user activation
func activatePost(c *gin.Context) {
	vm := models.ActivateVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	user, err := models.UsersDB.Activate(vm)
	if err != nil {
		abortWithError(c, http.StatusUnauthorized, err)
	}

	c.JSON(200, gin.H{"token": user.JWTToken})
}

//registerPost handles user registration
func registerPost(c *gin.Context) {
	vm := models.RegisterVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	vm.Email = helpers.NormalizeEmail(vm.Email)
	if err := checkmail.ValidateFormat(vm.Email); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if config.Settings.CheckEmails {
		if err := checkmail.ValidateHost(vm.Email); err != nil {
			abortWithError(c, http.StatusBadRequest, err)
			return
		}
	}

	user, err := models.UsersDB.Register(vm)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	go Email.SendUserRegistrationMessage(c, &user) //send email in background

	c.JSON(200, gin.H{"token": user.JWTToken})
}

//forgotPost handles password reset request
func forgotPost(c *gin.Context) {
	vm := models.ForgotVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	user, err := models.UsersDB.Forgot(vm)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	go Email.SendPasswordResetMessage(c, &user)

	c.JSON(http.StatusOK, gin.H{})
}

//resetPost handles password reset request
func resetPost(c *gin.Context) {
	vm := models.ResetVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	user, err := models.UsersDB.ResetPassword(vm)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	go Email.SendPasswordResetConfirmation(c, &user)

	c.JSON(200, gin.H{"token": user.JWTToken})
}
