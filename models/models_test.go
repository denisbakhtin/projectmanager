package models

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var userID uint64

const PASSWORD = "password"

func init() {
	os.Chdir("..")
	config.Initialize(gin.TestMode)
	InitializeDB()
	//create the required initial data for testing
	userID = getOrCreateUser().ID
	getOrCreateCategory()
	task := getOrCreateTask()
	getOrCreateComment(task.ID)
	getOrCreatePage()
	getOrCreateUnrelatedProject()
	getOrCreateUnrelatedSession()
	getOrCreateRelatedSession()
	getOrCreateSetting()
	getOrCreateTaskLog()
}

func TestInitializeDB(t *testing.T) {
	assert.NotNil(t, DB, "DB handler should not be nil")
	count := 0
	err := DB.Model(&UserGroup{}).Where([]int64{ADMIN, EDITOR, USER}).Count(&count).Error
	assert.Nil(t, err, "Error should be nil")
	assert.GreaterOrEqual(t, count, 3, "Not all required groups are present")

	err = DB.Model(&Setting{}).Where("code = ?", "site_name").First(&Setting{}).Error
	assert.Nil(t, err)
}

func getOrCreateCategory() Category {
	//this category should not have related projects or tasks
	cat, err := getUnrelatedCategory(userID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a category for testing")
	}
	if cat.ID != 0 {
		return cat
	}
	cat.Name = "Test Category"
	cat.UserID = userID
	if err := DB.Create(&cat).Error; err != nil {
		panic("Can't create a category for testing")
	}
	return cat
}

func getUnrelatedCategory(userID uint64) (Category, error) {
	cat := Category{}
	err := DB.
		Where("user_id = ? and NOT EXISTS(select null from projects where projects.category_id = categories.id) and NOT EXISTS(select null from tasks where tasks.category_id = categories.id)", userID).
		First(&cat).Error
	return cat, err
}

func getOrCreateTask() Task {
	task := Task{}
	err := DB.Where("user_id = ?", userID).First(&task).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a task for testing")
	}
	if task.ID != 0 {
		return task
	}
	project := getOrCreateUnrelatedProject()
	task.Name = "Test Task"
	task.UserID = userID
	task.ProjectID = project.ID
	if err := DB.Create(&task).Error; err != nil {
		panic("Can't create a task for testing")
	}
	return task
}

func getOrCreateComment(taskID uint64) Comment {
	comment := Comment{}
	err := DB.Where("user_id = ? and task_id = ?", userID, taskID).First(&comment).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a comment for testing")
	}
	if comment.ID != 0 {
		return comment
	}
	comment.Contents = "Test comment"
	comment.UserID = userID
	comment.TaskID = taskID
	if err := DB.Create(&comment).Error; err != nil {
		panic("Can't create a comment for testing")
	}
	return comment
}

func getOrCreatePage() Page {
	page := Page{}
	err := DB.First(&page).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a page for testing")
	}
	if page.ID != 0 {
		return page
	}
	page.Name = "Test page"
	if err := DB.Create(&page).Error; err != nil {
		panic("Can't create a page for testing")
	}
	return page
}

func getOrCreateUnrelatedProject() Project {
	project := Project{}
	err := DB.
		Where("user_id = ? and NOT EXISTS(select null from tasks where tasks.project_id = projects.id)", userID).
		First(&project).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a project for testing")
	}
	if project.ID != 0 {
		return project
	}
	project.Name = "Test project"
	project.UserID = userID
	if err := DB.Create(&project).Error; err != nil {
		panic("Can't create a project for testing")
	}
	return project
}

func getOrCreateUnrelatedSession() Session {
	session := Session{}
	err := DB.
		Where("user_id = ? and NOT EXISTS(select null from task_logs where task_logs.session_id = sessions.id)", userID).
		First(&session).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a session for testing")
	}
	if session.ID != 0 {
		return session
	}
	session.Contents = "Test session"
	session.UserID = userID
	if err := DB.Create(&session).Error; err != nil {
		panic("Can't create a session for testing")
	}
	return session
}

func getOrCreateRelatedSession() Session {
	session := Session{}
	err := DB.
		Where("user_id = ? and EXISTS(select null from task_logs where task_logs.session_id = sessions.id)", userID).
		First(&session).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a session for testing")
	}
	if session.ID != 0 {
		return session
	}
	session.Contents = "Test related session"
	session.UserID = userID
	session.TaskLogs = []TaskLog{
		{Minutes: 1},
	}
	if err := DB.Create(&session).Error; err != nil {
		panic("Can't create a related session for testing")
	}
	return session
}

func getOrCreateSetting() Setting {
	setting := Setting{}
	err := DB.First(&setting).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a setting for testing")
	}
	if setting.ID != 0 {
		return setting
	}
	setting.Code = "test_setting"
	setting.Value = "Test setting value"
	if err := DB.Create(&setting).Error; err != nil {
		panic("Can't create a related setting for testing")
	}
	return setting
}

func getOrCreateTaskLog() TaskLog {
	tl := TaskLog{}
	err := DB.Where("session_id = 0 and user_id = ?", userID).First(&tl).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic("Can't retreive a task log for testing")
	}
	if tl.ID != 0 {
		return tl
	}
	task := getOrCreateTask()
	tl.Minutes = 1
	tl.UserID = userID
	tl.TaskID = task.ID
	if err := DB.Create(&tl).Error; err != nil {
		panic("Can't create a task log for testing")
	}
	return tl
}

func getOrCreateUser() User {
	user := User{}
	err := DB.First(&user).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		panic(err)
	}
	if user.ID != 0 {
		return user
	}
	user = createUser("")

	return user
}

func createUser(password string) User {
	if len(password) == 0 {
		password = PASSWORD
	}
	now := time.Now().Nanosecond()
	phash := helpers.CreatePasswordHash(password)
	user := User{
		Name:         fmt.Sprintf("Test user %d", now),
		UserGroupID:  USER,
		Status:       ACTIVE,
		Email:        fmt.Sprintf("%d@email.com", now),
		PasswordHash: phash,
	}
	if err := DB.Create(&user).Error; err != nil {
		panic("Can't create a user for testing")
	}
	return user
}
