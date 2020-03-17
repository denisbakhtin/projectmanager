package controllers

import (
	"fmt"
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//sessionsGet handles get all sessions request
func searchGet(c *gin.Context) {
	query := c.Query("query")
	vm := models.SearchVM{}
	userID := currentUserID(c)

	models.DB.
		Where("user_id = ?", userID).
		Where("to_tsvector(name) @@ to_tsquery(?)", fmt.Sprintf("%s:*", query)).
		Order("id asc").Find(&vm.Categories)

	models.DB.Preload("Tasks").Preload("Category").
		Where("user_id = ?", userID).
		Where("to_tsvector(name || ' ' || description) @@ to_tsquery(?)", fmt.Sprintf("%s:*", query)).
		Order("archived asc, created_at asc").Find(&vm.Projects)

	q := models.DB.
		Where("user_id = ?", userID).
		Where("to_tsvector(name || ' ' || description) @@ to_tsquery(?)", fmt.Sprintf("%s:*", query)).
		Preload("Project").Preload("Comments").Preload("Category")
	q = q.Preload("TaskLogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("session_id = 0 and minutes > 0")
	})
	q.Order("tasks.completed asc, tasks.created_at asc").Find(&vm.Tasks)

	models.DB.
		Where("user_id = ?", userID).
		Where("to_tsvector(contents) @@ to_tsquery(?)", fmt.Sprintf("%s:*", query)).
		Order("id").Find(&vm.Comments)

	c.JSON(http.StatusOK, vm)
}
