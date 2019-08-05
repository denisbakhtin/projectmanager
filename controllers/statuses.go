package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//statusesGetHandler handles get all statuses request
func statusesGetHandler(c *gin.Context) {
	var statuses []models.Status
	models.DB.Order("order").Find(&statuses)
	c.JSON(http.StatusOK, statuses)
}

//statusGetHandler handles get status request
func statusGetHandler(c *gin.Context) {
	id := c.Param("id")
	status := models.Status{}
	models.DB.First(&status, id)
	if status.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}
	c.JSON(http.StatusOK, status)
}

//statusesPostHandler handles create status request
func statusesPostHandler(c *gin.Context) {
	status := models.Status{}
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//statusesPutHandler handles update status request
func statusesPutHandler(c *gin.Context) {
	//id := c.Param("id")
	status := models.Status{}
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//statusesDeleteHandler handles delete status request
func statusesDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	status := models.Status{}
	models.DB.First(&status, id)
	if status.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}
	if err := models.DB.Delete(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
