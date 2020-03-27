package controllers

import (
	"fmt"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/denisbakhtin/projectmanager/services"
	"github.com/gin-gonic/gin"
)

//loginPost handles user login
func loginPost(c *gin.Context) {
	vm := models.LoginVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	models.DB.Model(&models.User{}).Where("email = ?", helpers.NormalizeEmail(vm.Email)).First(&user)
	switch {
	case user.ID == 0 || !user.HasPassword(vm.Password):
		abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Wrong email or password"))
		return
	case user.Status == models.NOTACTIVE:
		abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Account requires activation"))
		return
	case user.Status == models.SUSPENDED:
		abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Account suspended"))
		return
	}

	if err := user.CreateJWTToken(); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{"token": user.Token})
}

//activatePost handles user activation
func activatePost(c *gin.Context) {
	vm := models.ActivateVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	models.DB.Where("token = ?", vm.Token).First(&user)
	switch {
	case user.ID == 0:
		abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Wrong activation token"))
		return
	case user.Status == models.SUSPENDED:
		abortWithError(c, http.StatusUnauthorized, fmt.Errorf("Account suspended"))
		return
	}
	//update user record
	user.Status = models.ACTIVE
	user.Token = ""
	if err := models.DB.Save(&user).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	if err := user.CreateJWTToken(); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{"token": user.Token})
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

	user := models.User{}
	models.DB.Where("email = ?", helpers.NormalizeEmail(vm.Email)).First(&user)
	if user.ID != 0 && user.Status != models.NOTACTIVE {
		abortWithError(c, http.StatusBadRequest, fmt.Errorf("This email already taken"))
		return
	}

	user.Name = vm.Name
	user.Email = vm.Email
	user.PasswordHash = helpers.CreatePasswordHash(vm.Password)
	user.UserGroupID = models.USER
	//user.Token = helpers.CreateSecureToken()
	user.Status = models.ACTIVE
	user.Token = ""

	//create new or update inactive account
	if err := models.DB.Save(&user).Error; err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	go services.SendUserRegistrationMessage(c, &user) //send email in background

	if err := user.CreateJWTToken(); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{"token": user.Token})
}
