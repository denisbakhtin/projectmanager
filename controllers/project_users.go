package controllers

/*
import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//projectUsersGet handles get all project users request
func projectUsersGet(c *gin.Context) {
	projectID := c.Param("project_id")
	user := models.User{}
	if u, exists := c.Get("user"); exists {
		user = u.(models.User)
	}
	project := models.Project{}
	models.DB.First(&project, projectID)
	if project.ID == 0 {
		c.JSON(http.StatusBadRequest, "Project not found")
		return
	}
	var pusers []models.ProjectUser
	models.DB.Preload("User").Preload("Role").Where("project_id = ?", projectID).Find(&pusers)
	if !user.BelongsToProjectUsers(pusers) {
		c.JSON(http.StatusBadRequest, "You do not belong to project users")
		return
	}
	c.JSON(http.StatusOK, pusers)
}

//projectUserGet handles get project user request
func projectUserGet(c *gin.Context) {
	projectID := c.Param("project_id")
	id := c.Param("id")
	puser := models.ProjectUser{}
	models.DB.Preload("User").Preload("Role").Where("project_id = ? and id = ?", projectID, id).First(&puser)
	if puser.ID == 0 {
		c.JSON(http.StatusNotFound, "Project user not found")
		return
	}
	c.JSON(http.StatusOK, puser)
}

//projectUsersPost handles create project user request
func projectUsersPost(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := models.User{}
	if u, exists := c.Get("user"); exists {
		user = u.(models.User)
	}
	puser := models.ProjectUser{}
	models.DB.Where("user_id = ?", user.ID).First(&puser)
	if puser.ID == 0 {
		c.JSON(http.StatusBadRequest, "Project user not found")
		return
	}
	task.ProjectUserID = puser.ID
	task.Completed = false
	if err := models.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectUsersPut handles update project user request
func projectUsersPut(c *gin.Context) {
	//id := c.Param("id")
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectUsersDelete handles delete project user request
func projectUsersDelete(c *gin.Context) {
	id := c.Param("id")
	task := models.Task{}
	models.DB.First(&task, id)
	if task.ID == 0 {
		c.JSON(http.StatusNotFound, "Task not found")
		return
	}
	if err := models.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
*/
