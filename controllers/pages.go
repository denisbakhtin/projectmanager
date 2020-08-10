package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//pagesGet handles get all pages request
func pagesGet(c *gin.Context) {
	pages, err := models.PagesDB.GetAll()
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, pages)
}

//pageGet handles get page request
func pageGet(c *gin.Context) {
	page, err := models.PagesDB.Get(c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Page"))
		} else {
			abortWithError(c, http.StatusInternalServerError, err)
		}
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
	page, err := models.PagesDB.Create(page)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, page)
}

//pagesPut handles update page request
func pagesPut(c *gin.Context) {
	page := models.Page{}
	if err := c.ShouldBindJSON(&page); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	page, err := models.PagesDB.Update(page)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, page)
}

//pagesDelete handles delete page request
func pagesDelete(c *gin.Context) {
	if err := models.PagesDB.Delete(c.Param("id")); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//pagesGetHTML handles get html page request
func pagesGetHTML(c *gin.Context) {
	page, _ := models.PagesDB.Get(c.Param("id"))
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
