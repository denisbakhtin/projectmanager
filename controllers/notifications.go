package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//notificationsGet handles get all notifications request
func notificationsGet(c *gin.Context) {
	var notifications []models.Notification
	if user, ok := c.Get("user"); ok {
		if u, ok := user.(models.User); ok {
			models.DB.Where("user_id = ?", u.ID).Find(&notifications)
		}
	}
	c.JSON(http.StatusOK, notifications)
}

//notificationsDelete handles delete notification request
func notificationsDelete(c *gin.Context) {
	id := c.Param("id")
	notification := models.Notification{}
	if user, ok := c.Get("user"); ok {
		if u, ok := user.(models.User); ok {
			models.DB.Where("user_id = ?", u.ID).First(&notification, id)
			models.DB.Delete(&notification)
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}
