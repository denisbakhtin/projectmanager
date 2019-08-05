package controllers

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/denisbakhtin/projectmanager/services"
	"github.com/denisbakhtin/projectmanager/viewmodels"
	"github.com/gin-gonic/gin"
)

//loginPostHandler handles user login
func loginPostHandler(c *gin.Context) {
	vm := viewmodels.LoginVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}
	models.DB.Model(&models.User{}).Where("email = ?", helpers.NormalizeEmail(vm.Email)).First(&user)
	switch {
	case user.ID == 0 || !user.HasPassword(vm.Password):
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong email or password"})
		return
	case user.Status == models.NOTACTIVE:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account requires activation"})
		return
	case user.Status == models.SUSPENDED:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account suspended"})
		return
	}

	if err := user.CreateJWTToken(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": user.Token})
}

//activatePostHandler handles user activation
func activatePostHandler(c *gin.Context) {
	vm := viewmodels.ActivateVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}
	models.DB.Where("token = ?", vm.Token).First(&user)
	switch {
	case user.ID == 0:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong activation token"})
		return
	case user.Status == models.SUSPENDED:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account suspended"})
		return
	}
	//update user record
	user.Status = models.ACTIVE
	user.Token = ""
	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := user.CreateJWTToken(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": user.Token})
}

//registerPostHandler handles user registration
func registerPostHandler(c *gin.Context) {
	vm := viewmodels.RegisterVM{}
	if err := c.ShouldBindJSON(&vm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vm.Email = helpers.NormalizeEmail(vm.Email)
	if err := checkmail.ValidateFormat(vm.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if config.Settings.CheckEmails {
		if err := checkmail.ValidateHost(vm.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	user := models.User{}
	models.DB.Where("email = ?", helpers.NormalizeEmail(vm.Email)).First(&user)
	if user.ID != 0 && user.Status != models.NOTACTIVE {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email already taken"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go services.SendUserRegistrationMessage(c, &user) //send email in background

	if err := user.CreateJWTToken(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": user.Token})
}
