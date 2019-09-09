package controllers

import (
	"net/http"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//tasksGet handles get all tasks request
func tasksGet(c *gin.Context) {
	var tasks []models.Task
	models.DB.Preload("ProjectUser").Preload("ProjectUser.User").Preload("Project").Preload("TaskStep").Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

//taskGet handles get task request
func taskGet(c *gin.Context) {
	id := c.Param("id")
	task := models.Task{}
	models.DB.Preload("AttachedFiles").First(&task, id)
	if task.ID == 0 {
		c.JSON(http.StatusNotFound, "Task not found")
		return
	}
	c.JSON(http.StatusOK, task)
}

//taskNewGet handles get new task request
func taskNewGet(c *gin.Context) {
	vm := models.EditTaskVM{}
	userID := currentUserID(c)
	//TODO: filter projects by owner and assigned users
	models.DB.Where("owner_id = ?", userID).Find(&vm.Projects)
	models.DB.Order("ord asc").Find(&vm.TaskSteps)
	if len(vm.TaskSteps) > 0 {
		vm.Task.TaskStepID = vm.TaskSteps[0].ID
	}
	if len(vm.Projects) > 0 {
		vm.Task.ProjectID = vm.Projects[0].ID
	}
	c.JSON(http.StatusOK, vm)
}

//taskEditGet handles get edit task request
func taskEditGet(c *gin.Context) {
	vm := models.EditTaskVM{}
	userID := currentUserID(c)
	//TODO: filter projects by owner and assigned users
	models.DB.Where("owner_id = ?", userID).Find(&vm.Projects)
	models.DB.Order("ord asc").Find(&vm.TaskSteps)
	models.DB.Preload("AttachedFiles").Preload("Project").Preload("ProjectUser").Preload("TaskStep").First(&vm.Task, c.Param("id"))
	c.JSON(http.StatusOK, vm)
}

//tasksPost handles create role request
func tasksPost(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := models.User{}
	if u, exists := c.Get("user"); exists {
		user = u.(models.User)
	}
	puser := models.ProjectUser{}
	models.DB.Where("user_id = ?", user.ID).First(&puser)
	if puser.ID == 0 {
		c.JSON(http.StatusBadRequest, "Project user not found")
		return
	}
	task.ProjectUserID = puser.ID
	step := models.TaskStep{}
	models.DB.Order("order").First(&step)
	task.TaskStepID = step.ID
	if err := models.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksPut handles update task request
func tasksPut(c *gin.Context) {
	//id := c.Param("id")
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

//tasksDelete handles delete role request
func tasksDelete(c *gin.Context) {
	id := c.Param("id")
	task := models.Task{}
	models.DB.First(&task, id)
	if task.ID == 0 {
		c.JSON(http.StatusNotFound, "Task not found")
		return
	}
	if err := models.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
