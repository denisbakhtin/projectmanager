package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//rolesGet handles get all roles request
func rolesGet(c *gin.Context) {
	var roles []models.Role
	models.DB.Find(&roles)
	c.JSON(http.StatusOK, roles)
}

//roleGet handles get role request
func roleGet(c *gin.Context) {
	id := c.Param("id")
	role := models.Role{}
	models.DB.First(&role, id)
	if role.ID == 0 {
		c.JSON(http.StatusNotFound, "Role not found")
		return
	}
	c.JSON(http.StatusOK, role)
}

//rolesPost handles create role request
func rolesPost(c *gin.Context) {
	role := models.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//rolesPut handles update role request
func rolesPut(c *gin.Context) {
	//id := c.Param("id")
	role := models.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//rolesDelete handles delete role request
func rolesDelete(c *gin.Context) {
	id := c.Param("id")
	role := models.Role{}
	models.DB.First(&role, id)
	if role.ID == 0 {
		c.JSON(http.StatusNotFound, "Role not found")
		return
	}
	if err := models.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
