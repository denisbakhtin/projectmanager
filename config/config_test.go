package config

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("..")
}

func TestInitialize(t *testing.T) {
	Initialize(gin.TestMode)

	assert.NotNil(t, LogFile, "Log file should not be empty")
	defer LogFile.Close()

	assert.Equal(t, UploadPathURL, "/public/uploads", "Wrong uploads url prefix")

	wd, _ := filepath.Abs("")
	assert.Equal(t, wd, AppDir, "Wrong working directory")
	assert.Equal(t, UploadPath, path.Join(wd, "public", "uploads"), "Wrong uploads dir path")

	assert.NotEmpty(t, SettingsAll, "Settings should not be empty")
	assert.NotEqual(t, SettingsAll.Test, SettingsAll.Production, "Test and production config should not match")
	assert.NotEqual(t, SettingsAll.Test, SettingsAll.Development, "Test and development config should not match")
	assert.Equal(t, SettingsAll.Test, Settings, "Should load test config")
	assert.NotEmpty(t, Settings, "Settings should not be empty")
}
