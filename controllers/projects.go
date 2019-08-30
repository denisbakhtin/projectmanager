package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//projectsGet handles get all projects request
func projectsGet(c *gin.Context) {
	var projects []models.Project
	models.DB.Preload("Owner").Preload("Status").Find(&projects)
	c.JSON(http.StatusOK, projects)
}

//projectGet handles get project request
func projectGet(c *gin.Context) {
	id := c.Param("id")
	project := models.Project{}
	models.DB.Preload("ProjectUsers").Preload("ProjectUsers.Role").Preload("ProjectUsers.User").Preload("AttachedFiles").Preload("Owner").Preload("Status").Preload("Tasks").First(&project, id)
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	c.JSON(http.StatusOK, project)
}

//projectsPost handles create role request
func projectsPost(c *gin.Context) {
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}
	if u, exists := c.Get("user"); exists {
		user = u.(models.User)
	}
	project.OwnerID = user.ID
	if err := models.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectsPut handles update project request
func projectsPut(c *gin.Context) {
	//id := c.Param("id")
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project.Status = models.Status{} //prevent gorm from taking its id instead of project.StatusID
	if err := models.DB.Omit("owner_id").Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//delete removed users
	userIds := []uint64{0} //add atleast one non-existent id for query to work :)
	for i := 0; i < len(project.ProjectUsers); i++ {
		userIds = append(userIds, project.ProjectUsers[i].ID)
	}
	if err := models.DB.Where("project_id = ? and id not in (?)", project.ID, userIds).Delete(models.ProjectUser{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//delete removed files
	fileIds := []uint64{0} //add atleast one non-existent id for query to work :)
	for i := 0; i < len(project.AttachedFiles); i++ {
		fileIds = append(fileIds, project.AttachedFiles[i].ID)
	}
	if err := models.DB.Where("owner_type = ? and owner_id = ? and id not in (?)", "projects", project.ID, fileIds).Delete(models.AttachedFile{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectsDelete handles delete role request
func projectsDelete(c *gin.Context) {
	id := c.Param("id")
	project := models.Project{}
	models.DB.First(&project, id)
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	if err := models.DB.Delete(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
