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
	projects, err := models.ProjectsDB.GetAll(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, projects)
}

//projectsFavoriteGet handles get all favorite projects request
func projectsFavoriteGet(c *gin.Context) {
	projects, err := models.ProjectsDB.GetAllFavorite(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, projects)
}

//projectGet handles get project request
func projectGet(c *gin.Context) {
	project, err := models.ProjectsDB.Get(currentUserID(c), c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Page"))
		} else {
			abortWithError(c, http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, project)
}

//projectNewGet handles get new project request
func projectNewGet(c *gin.Context) {
	vm, err := models.ProjectsDB.GetNew(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}

//projectEditGet handles edit project request
func projectEditGet(c *gin.Context) {
	vm, err := models.ProjectsDB.GetEdit(currentUserID(c), c.Param("id"))
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}

//projectsPost handles create role request
func projectsPost(c *gin.Context) {
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	if _, err := models.ProjectsDB.Create(currentUserID(c), project); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectsPut handles update project request
func projectsPut(c *gin.Context) {
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.ProjectsDB.Update(currentUserID(c), project); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectArchive handles archive project request
func projectArchive(c *gin.Context) {
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.ProjectsDB.ToggleArchive(currentUserID(c), project); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectFavorite handles toggling project favorite status request
func projectFavorite(c *gin.Context) {
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.ProjectsDB.ToggleFavorite(currentUserID(c), project); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//projectsDelete handles delete role request
func projectsDelete(c *gin.Context) {
	if err := models.ProjectsDB.Delete(currentUserID(c), c.Param("id")); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//projectsSummaryGet handles get projects statistics request
func projectsSummaryGet(c *gin.Context) {
	vm, err := models.ProjectsDB.Summary(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}
