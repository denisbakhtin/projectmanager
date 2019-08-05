package controllers

import (
	"github.com/gin-gonic/gin"
)

//adminHandler handles admin's root html page request
func adminHandler(c *gin.Context) {
	c.HTML(200, "admin/dashboard.tmpl", gin.H{
		"Title": "Project Manager welcomes you!",
	})
}
