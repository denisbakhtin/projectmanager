package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//projectsGet handles get all projects request
func projectsGet(c *gin.Context) {
	var projects []models.Project
	models.DB.Preload("Tasks").Preload("Category").
		Where("user_id = ?", currentUserID(c)).Order("archived asc, created_at asc").
		Find(&projects)
	c.JSON(http.StatusOK, projects)
}

//projectsFavoriteGet handles get all favorite projects request
func projectsFavoriteGet(c *gin.Context) {
	var projects []models.Project
	models.DB.Select("id, name").Where("user_id = ? and favorite = true", currentUserID(c)).Order("created_at asc").Find(&projects)
	c.JSON(http.StatusOK, projects)
}

//projectGet handles get project request
func projectGet(c *gin.Context) {
	id := c.Param("id")
	project := models.Project{}
	query := models.DB.Where("user_id = ?", currentUserID(c)).
		Preload("AttachedFiles").Preload("Category")
	query = query.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Order("tasks.completed asc, created_at asc")
	})
	query = query.Preload("Tasks.Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at asc")
	})

	query.First(&project, id)
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, helpers.NotFoundOrOwned("Project"))
		return
	}
	c.JSON(http.StatusOK, project)
}

//projectNewGet handles get new project request
func projectNewGet(c *gin.Context) {
	vm := models.EditProjectVM{}
	userID := currentUserID(c)
	models.DB.Where("user_id = ?", userID).Find(&vm.Categories)
	c.JSON(http.StatusOK, vm)
}

//projectEditGet handles edit project request
func projectEditGet(c *gin.Context) {
	vm := models.EditProjectVM{}
	userID := currentUserID(c)
	models.DB.Where("user_id = ?", userID).Find(&vm.Categories)
	models.DB.Preload("AttachedFiles").Preload("Tasks").First(&vm.Project, c.Param("id"))
	c.JSON(http.StatusOK, vm)
}

//projectsPost handles create role request
func projectsPost(c *gin.Context) {
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	project.UserID = currentUserID(c)
	if err := models.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectsPut handles update project request
func projectsPut(c *gin.Context) {
	//id := c.Param("id")
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	project.UserID = currentUserID(c)
	if err := models.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectArchive handles archive project request
func projectArchive(c *gin.Context) {
	//id := c.Param("id")
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//update column without hooks
	if err := models.DB.Model(&project).UpdateColumn("archived", project.Archived).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectFavorite handles toggling project favorite status request
func projectFavorite(c *gin.Context) {
	//id := c.Param("id")
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//update column without hooks
	if err := models.DB.Model(&project).UpdateColumn("favorite", project.Favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectsDelete handles delete role request
func projectsDelete(c *gin.Context) {
	id := c.Param("id")
	project := models.Project{}
	models.DB.Where("user_id = ?", currentUserID(c)).First(&project, id)
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, helpers.NotFoundOrOwned("Project"))
		return
	}
	if err := models.DB.Delete(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
