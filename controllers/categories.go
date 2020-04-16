package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//categoriesGet handles get all categories request
func categoriesGet(c *gin.Context) {
	categories, err := models.CategoriesDB.GetAll(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, categories)
}

//categoryGet handles get category request
func categoryGet(c *gin.Context) {
	category, err := models.CategoriesDB.Get(currentUserID(c), c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Category"))
		} else {
			abortWithError(c, http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, category)
}

//categoriesPost handles create category request
func categoriesPost(c *gin.Context) {
	category := models.Category{}
	if err := c.ShouldBindJSON(&category); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.CategoriesDB.Create(currentUserID(c), category); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//categoriesPut handles update category request
func categoriesPut(c *gin.Context) {
	category := models.Category{}
	if err := c.ShouldBindJSON(&category); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.CategoriesDB.Update(currentUserID(c), category); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//categoriesDelete handles delete category request
func categoriesDelete(c *gin.Context) {
	if err := models.CategoriesDB.Delete(currentUserID(c), c.Param("id")); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//categoriesSummaryGet handles get categories statistics request
func categoriesSummaryGet(c *gin.Context) {
	vm, err := models.CategoriesDB.Summary(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}
