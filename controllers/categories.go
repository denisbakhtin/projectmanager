package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//categoriesGet handles get all categories request
func categoriesGet(c *gin.Context) {
	var categories []models.Category
	models.DB.Where("user_id = ?", currentUserID(c)).Order("id asc").Find(&categories)
	c.JSON(http.StatusOK, categories)
}

//categoryGet handles get category request
func categoryGet(c *gin.Context) {
	id := c.Param("id")
	category := models.Category{}
	models.DB.Where("user_id = ?", currentUserID(c)).Preload("Tasks").
		Preload("Tasks.Comments").Preload("Projects").Preload("Projects.Tasks").First(&category, id)
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, helpers.NotFoundOrOwned("Category"))
		return
	}
	c.JSON(http.StatusOK, category)
}

//categoriesPost handles create category request
func categoriesPost(c *gin.Context) {
	category := models.Category{}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	category.UserID = currentUserID(c)
	if err := models.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//categoriesPut handles update category request
func categoriesPut(c *gin.Context) {
	//id := c.Param("id")
	category := models.Category{}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	category.UserID = currentUserID(c)
	if err := models.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//categoriesDelete handles delete category request
func categoriesDelete(c *gin.Context) {
	id := c.Param("id")
	category := models.Category{}
	userID := currentUserID(c)
	models.DB.Where("user_id = ?", userID).First(&category, id)
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, helpers.NotFoundOrOwned("Category"))
		return
	}
	if err := models.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//categoriesSummaryGet handles get categories statistics request
func categoriesSummaryGet(c *gin.Context) {
	vm := models.CategoriesSummaryVM{}
	userID := currentUserID(c)
	if err := models.DB.Model(models.Category{}).Where("user_id = ?", userID).Count(&vm.Count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vm)
}
