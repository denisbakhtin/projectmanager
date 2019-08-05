package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//rolesGetHandler handles get all roles request
func rolesGetHandler(c *gin.Context) {
	var roles []models.Role
	models.DB.Find(&roles)
	c.JSON(http.StatusOK, roles)
}

//roleGetHandler handles get role request
func roleGetHandler(c *gin.Context) {
	id := c.Param("id")
	role := models.Role{}
	models.DB.First(&role, id)
	if role.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

//rolesPostHandler handles create role request
func rolesPostHandler(c *gin.Context) {
	role := models.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//rolesPutHandler handles update role request
func rolesPutHandler(c *gin.Context) {
	//id := c.Param("id")
	role := models.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//rolesDeleteHandler handles delete role request
func rolesDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	role := models.Role{}
	models.DB.First(&role, id)
	if role.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	if err := models.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
