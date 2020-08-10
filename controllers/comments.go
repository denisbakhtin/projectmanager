package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//commentsGet handles GET all comments request
func commentsGet(c *gin.Context) {
	comments, err := models.CommentsDB.GetAll(currentUserID(c), c.Param("task_id"))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, comments)
}

//commentGet handles get comment request
func commentGet(c *gin.Context) {
	comment, err := models.CommentsDB.Get(currentUserID(c), c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Comment"))
		} else {
			abortWithError(c, http.StatusInternalServerError, err)
		}
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
	com, err := models.CommentsDB.Create(currentUserID(c), comment)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, com)
}

//commentsPut handles update comment request
func commentsPut(c *gin.Context) {
	comment := models.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	com, err := models.CommentsDB.Update(currentUserID(c), comment)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, com)
}

//commentsDelete handles delete comment request
func commentsDelete(c *gin.Context) {
	if err := models.CommentsDB.Delete(currentUserID(c), c.Param("id")); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
