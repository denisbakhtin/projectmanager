package controllers

import (
	"log"
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//pagesGetHandler handles get all pages request
func pagesGetHandler(c *gin.Context) {
	var pages []models.Page
	models.DB.Find(&pages)
	c.JSON(http.StatusOK, pages)
}

//pageGetHandler handles get page request
func pageGetHandler(c *gin.Context) {
	id := c.Param("id")
	page := models.Page{}
	models.DB.First(&page, id)
	if page.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}
	c.JSON(http.StatusOK, page)
}

//pagesPostHandler handles create page request
func pagesPostHandler(c *gin.Context) {
	page := models.Page{}
	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//pagesPutHandler handles update page request
func pagesPutHandler(c *gin.Context) {
	//id := c.Param("id")
	page := models.Page{}
	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//pagesDeleteHandler handles delete page request
func pagesDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	page := models.Page{}
	models.DB.First(&page, id)
	if page.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}
	if err := models.DB.Delete(&page).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//pagesGetHTMLHandler handles get html page request
func pagesGetHTMLHandler(c *gin.Context) {
	id := c.Param("id")
	log.Println("============================== Page id = ", id)

	page := models.Page{}
	models.DB.First(&page, id)
	if page.ID == 0 || !page.Published {
		c.HTML(http.StatusNotFound, "errors/404.tmpl", gin.H{"error": "Requested page not found"})
		return
	}
	c.HTML(http.StatusOK, "pages/page.tmpl", gin.H{
		"Title": page.Name,
		"Page":  page,
	})
}
