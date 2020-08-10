package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//settingsGet handles GET all settings request
func settingsGet(c *gin.Context) {
	settings, err := models.SettingsDB.GetAll()
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, settings)
}

//settingGet handles get setting request
func settingGet(c *gin.Context) {
	setting, err := models.SettingsDB.Get(c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Setting"))
		} else {
			abortWithError(c, http.StatusInternalServerError, err)
		}
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
	setting, err := models.SettingsDB.Create(setting)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, setting)
}

//settingsPut handles update setting request
func settingsPut(c *gin.Context) {
	setting := models.Setting{}
	if err := c.ShouldBindJSON(&setting); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	setting, err := models.SettingsDB.Update(setting)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, setting)
}

//settingsDelete handles delete setting request
func settingsDelete(c *gin.Context) {
	if err := models.SettingsDB.Delete(c.Param("id")); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
