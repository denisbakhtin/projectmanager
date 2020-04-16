package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//tasksGet handles get all tasks request
func tasksGet(c *gin.Context) {
	tasks, err := models.TasksDB.GetAll(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

//taskGet handles get task request
func taskGet(c *gin.Context) {
	task, err := models.TasksDB.Get(currentUserID(c), c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			abortWithError(c, http.StatusNotFound, helpers.NotFoundOrOwnedError("Task"))
		} else {
			abortWithError(c, http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, task)
}

//taskNewGet handles get new task request
func taskNewGet(c *gin.Context) {
	projectID := helpers.AtoUint64(c.Query("project_id"))
	vm, err := models.TasksDB.GetNew(currentUserID(c), projectID)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, vm)
}

//taskEditGet handles get edit task request
func taskEditGet(c *gin.Context) {
	vm, err := models.TasksDB.GetEdit(currentUserID(c), c.Param("id"))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, vm)
}

//tasksPost handles create task request
func tasksPost(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.TasksDB.Create(currentUserID(c), task); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksPut handles update task request
func tasksPut(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	if _, err := models.TasksDB.Update(currentUserID(c), task); err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksDelete handles delete task request
func tasksDelete(c *gin.Context) {
	if err := models.TasksDB.Delete(currentUserID(c), c.Param("id")); err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//tasksSummaryGet handles get tasks statistics request
func tasksSummaryGet(c *gin.Context) {
	vm, err := models.TasksDB.Summary(currentUserID(c))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, vm)
}
