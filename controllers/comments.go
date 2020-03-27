package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//commentsGet handles GET all comments request
func commentsGet(c *gin.Context) {
	var comments []models.Comment
	models.DB.Where("user_id = ? and task_id = ?", currentUserID(c), c.Param("task_id")).Order("id").Find(&comments)
	c.JSON(http.StatusOK, comments)
}

//commentGet handles get comment request
func commentGet(c *gin.Context) {
	id := c.Param("id")
	comment := models.Comment{}
	models.DB.Where("user_id = ?", currentUserID(c)).First(&comment, id)
	if comment.ID == 0 {
		abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Comment"))
		return
	}
	c.JSON(http.StatusOK, comment)
}

//commentsPost handles create comment request
func commentsPost(c *gin.Context) {
	comment := models.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	comment.UserID = currentUserID(c)
	if err := models.DB.Create(&comment).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//commentsPut handles update comment request
func commentsPut(c *gin.Context) {
	comment := models.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	userID := currentUserID(c)
	comment.UserID = userID
	if err := models.DB.Save(&comment).Error; err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//commentsDelete handles delete comment request
func commentsDelete(c *gin.Context) {
	id := c.Param("id")
	comment := models.Comment{}
	models.DB.Where("user_id = ?", currentUserID(c)).First(&comment, id)
	if comment.ID == 0 {
		abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Comment"))
		return
	}
	if err := models.DB.Delete(&comment).Error; err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
