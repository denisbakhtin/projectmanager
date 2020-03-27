package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

func uploadsPost(c *gin.Context) {
	uploader := c.Param("uploader") // form or ckeditor
	file, _ := c.FormFile("upload")
	ext := filepath.Ext(file.Filename)
	if !isExtensionAllowed(ext) || !isValidFilename(file.Filename) {
		//TODO: extend the list? Or just prohibit executables
		abortWithError(c, http.StatusBadRequest, fmt.Errorf("Unsupported file extension"))
		return
	}
	user := c.MustGet("user").(models.User)
	now := time.Now()
	subDir := path.Join(fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day()), fmt.Sprintf("user_%d", user.ID))
	uploadsDir := path.Join(config.UploadPath, subDir)
	filename, err := helpers.GetUniqueFilename(uploadsDir, file.Filename)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := c.SaveUploadedFile(file, path.Join(uploadsDir, filename)); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	fileurl := config.UploadPathURL + "/" + subDir + "/" + filename
	if strings.Compare(uploader, "ckeditor") == 0 {
		ckfunc := c.PostForm("CKEditorFuncNum")
		c.String(http.StatusOK, "<script>window.parent.CKEDITOR.tools.callFunction("+ckfunc+", \""+fileurl+"\");</script>")
		return
	}
	c.JSON(http.StatusOK, models.AttachedFile{
		Name:     filename,
		FilePath: path.Join(uploadsDir, filename),
		URL:      fileurl,
	})
}

//fileIsSafeToSave checks if file extension is acceptable
func isExtensionAllowed(fileExt string) bool {
	validExtPattern := "(?i)^\\.doc[x]?|\\.jp[e]?g|\\.bmp|\\.zip|\\.rar|\\.gz|\\.gzip|\\.gif|\\.txt|\\.pdf|\\.png$" //(?i) == case insensitive
	if matched, _ := regexp.MatchString(validExtPattern, fileExt); matched {
		return true
	}
	return false
}

//isValidForRemoval checks if filename excludes certain dangerous symbols
func isValidFilename(fileName string) bool {
	switch {
	case fileName == "":
		return false
	case strings.Contains(fileName, "/") || strings.Contains(fileName, "\\"):
		return false //no change dir plz
	case strings.Contains(fileName, "*") || strings.Contains(fileName, "?"):
		return false //no masks
	default:
		return true
	}
}
