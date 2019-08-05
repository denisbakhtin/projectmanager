package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//usersGetHandler handles get all users request
//TODO: allow access only to admins via middleware
func usersGetHandler(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

//userGetHandler handles get user request
func userGetHandler(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	models.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

//usersPutHandler handles user status update
func usersPutHandler(c *gin.Context) {
	//id := c.Param("id")
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("id = ?", user.ID).Update("status", user.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
