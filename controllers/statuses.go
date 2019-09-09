package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//statusesGet handles get all statuses request
func statusesGet(c *gin.Context) {
	var statuses []models.Status
	models.DB.Order("order").Find(&statuses)
	c.JSON(http.StatusOK, statuses)
}

//statusGet handles get status request
func statusGet(c *gin.Context) {
	id := c.Param("id")
	status := models.Status{}
	models.DB.First(&status, id)
	if status.ID == 0 {
		c.JSON(http.StatusNotFound, "Status not found")
		return
	}
	c.JSON(http.StatusOK, status)
}

//statusesPost handles create status request
func statusesPost(c *gin.Context) {
	status := models.Status{}
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Create(&status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//statusesPut handles update status request
func statusesPut(c *gin.Context) {
	//id := c.Param("id")
	status := models.Status{}
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Save(&status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//statusesDelete handles delete status request
func statusesDelete(c *gin.Context) {
	id := c.Param("id")
	status := models.Status{}
	models.DB.First(&status, id)
	if status.ID == 0 {
		c.JSON(http.StatusNotFound, "Status not found")
		return
	}
	if err := models.DB.Delete(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
