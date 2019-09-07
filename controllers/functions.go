package controllers

import (
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//currentUserID returns authenticated user ID
func currentUserID(c *gin.Context) uint64 {
	if u, exists := c.Get("user"); exists {
		if user, ok := u.(models.User); ok {
			return user.ID
		}
	}
	return 0
}
