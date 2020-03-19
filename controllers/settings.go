package controllers

import (
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
		c.JSON(http.StatusNotFound, "Setting not found")
		return
	}
	c.JSON(http.StatusOK, setting)
}

//settingsPost handles create setting request
func settingsPost(c *gin.Context) {
	setting := models.Setting{}
	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Create(&setting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//settingsPut handles update setting request
func settingsPut(c *gin.Context) {
	//id := c.Param("id")
	setting := models.Setting{}
	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Save(&setting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusNotFound, "Setting not found")
		return
	}
	if err := models.DB.Delete(&setting).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}