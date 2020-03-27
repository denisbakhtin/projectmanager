package controllers

import (
	"fmt"
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//usersGet handles get all users request
func usersGet(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

//userGet handles get user request
func userGet(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	models.DB.First(&user, id)
	if user.ID == 0 {
		abortWithError(c, http.StatusNotFound, fmt.Errorf("User not found"))
		return
	}
	c.JSON(http.StatusOK, user)
}

//usersPut handles user status update
func usersPut(c *gin.Context) {
	//id := c.Param("id")
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := models.DB.Where("id = ?", user.ID).Update("status", user.Status).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//usersSummaryGet handles get users statistics request
func usersSummaryGet(c *gin.Context) {
	vm := models.UsersSummaryVM{}
	if err := models.DB.Model(models.User{}).Count(&vm.Count).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}
