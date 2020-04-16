package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//sessionsGet handles get all sessions request
func searchGet(c *gin.Context) {
	vm, err := models.SearchDB.Search(currentUserID(c), c.Query("query"))

	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, vm)
}
