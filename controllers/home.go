package controllers

import (
	"github.com/denisbakhtin/projectmanager/config"
	"github.com/gin-gonic/gin"
)

//home handles home page html request
func home(c *gin.Context) {
	c.HTML(200, "home/home", gin.H{
		"Title": config.Settings.ProjectName + " welcomes you!",
	})
}
