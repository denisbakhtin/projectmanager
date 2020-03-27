package controllers

import (
	"fmt"
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//pagesGet handles get all pages request
func pagesGet(c *gin.Context) {
	var pages []models.Page
	models.DB.Order("published desc, id asc").Find(&pages)
	c.JSON(http.StatusOK, pages)
}

//pageGet handles get page request
func pageGet(c *gin.Context) {
	id := c.Param("id")
	page := models.Page{}
	models.DB.First(&page, id)
	if page.ID == 0 {
		abortWithError(c, http.StatusNotFound, fmt.Errorf("Page not found"))
		return
	}
	c.JSON(http.StatusOK, page)
}

//pagesPost handles create page request
func pagesPost(c *gin.Context) {
	page := models.Page{}
	if err := c.ShouldBindJSON(&page); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := models.DB.Create(&page).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//pagesPut handles update page request
func pagesPut(c *gin.Context) {
	//id := c.Param("id")
	page := models.Page{}
	if err := c.ShouldBindJSON(&page); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := models.DB.Save(&page).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//pagesDelete handles delete page request
func pagesDelete(c *gin.Context) {
	id := c.Param("id")
	page := models.Page{}
	models.DB.First(&page, id)
	if page.ID == 0 {
		abortWithError(c, http.StatusNotFound, fmt.Errorf("Page not found"))
		return
	}
	if err := models.DB.Delete(&page).Error; err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//pagesGetHTML handles get html page request
func pagesGetHTML(c *gin.Context) {
	id := c.Param("id")

	page := models.Page{}
	models.DB.First(&page, id)
	if page.ID == 0 || !page.Published {
		c.HTML(http.StatusNotFound, "errors/404", "Requested page not found")
		return
	}
	c.HTML(http.StatusOK, "pages/page", gin.H{
		"Title":           page.Name,
		"MetaKeywords":    page.MetaKeywords,
		"MetaDescription": page.MetaDescription,
		"Page":            &page,
	})
}
