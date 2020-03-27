package controllers

import (
	"fmt"
	"net/http"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//settingsGet handles GET all settings request
func settingsGet(c *gin.Context) {
	var settings []models.Setting
	models.DB.Order("id").Find(&settings)
	//append some settings from config.yml
	settings = append(settings, models.Setting{Code: "site_name", Value: config.Settings.ProjectName})
	c.JSON(http.StatusOK, settings)
}

//settingGet handles get setting request
func settingGet(c *gin.Context) {
	id := c.Param("id")
	setting := models.Setting{}
	models.DB.First(&setting, id)
	if setting.ID == 0 {
		abortWithError(c, http.StatusNotFound, fmt.Errorf("Setting not found"))
		return
	}
	c.JSON(http.StatusOK, setting)
}

//settingsPost handles create setting request
func settingsPost(c *gin.Context) {
	setting := models.Setting{}
	if err := c.ShouldBindJSON(&setting); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := models.DB.Create(&setting).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//settingsPut handles update setting request
func settingsPut(c *gin.Context) {
	//id := c.Param("id")
	setting := models.Setting{}
	if err := c.ShouldBindJSON(&setting); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := models.DB.Save(&setting).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//settingsDelete handles delete setting request
func settingsDelete(c *gin.Context) {
	id := c.Param("id")
	setting := models.Setting{}
	models.DB.First(&setting, id)
	if setting.ID == 0 {
		abortWithError(c, http.StatusNotFound, fmt.Errorf("Setting not found"))
		return
	}
	if err := models.DB.Delete(&setting).Error; err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
