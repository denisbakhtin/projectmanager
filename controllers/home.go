package controllers

import (
	"github.com/denisbakhtin/projectmanager/config"
	"github.com/gin-gonic/gin"
)

//homeHandler handles home page html request
func homeHandler(c *gin.Context) {
	c.HTML(200, "home/home.tmpl", gin.H{
		"Title": config.Settings.ProjectName + " welcomes you!",
	})
}
