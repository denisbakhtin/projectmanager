package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/controllers/mocks"
	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

var (
	router              *gin.Engine
	server              *httptest.Server
	usersDBMock         mocks.UsersDBMock
	categoriesDBMock    mocks.CategoriesDBMock
	tasksDBMock         mocks.TasksDBMock
	sessionsDBMock      mocks.SessionsDBMock
	commentsDBMock      mocks.CommentsDBMock
	pagesDBMock         mocks.PagesDBMock
	projectsDBMock      mocks.ProjectsDBMock
	userActivationToken string
	emailMock           mocks.EmailMock
	authenticatedUser   models.User
	authenticatedAdmin  models.User
)

func init() {
	os.Chdir("..")
	config.Initialize(gin.TestMode)
	models.InitializeDB()
	setupTestServer()
	seedUsersDB()
	authenticatedUser = getAuthenticatedUser()
	authenticatedAdmin = getAuthenticatedAdmin()
	seedCategoriesDB()
	seedTasksDB()
	seedSessionsDB()
	seedCommentsDB()
	seedPagesDB()
	seedProjectsDB()
	models.UsersDB = &usersDBMock
	models.CategoriesDB = &categoriesDBMock
	models.TaskLogsDB = &mocks.TaskLogsDBMock{}
	models.TasksDB = &tasksDBMock
	models.SessionsDB = &sessionsDBMock
	models.CommentsDB = &commentsDBMock
	models.PagesDB = &pagesDBMock
	models.ProjectsDB = &projectsDBMock
	Email = &emailMock
}

func setupTestServer() {
	router := setupRouter()
	if router == nil {
		panic("Nil router")
	}
	server = httptest.NewServer(router)
}

func seedUsersDB() {
	usersDBMock.Users = make([]models.User, 10)
	for i := 0; i < 10; i++ {
		usersDBMock.Users[i] = models.User{
			ID:           uint64(i),
			Name:         fmt.Sprintf("User%d", i),
			Email:        fmt.Sprintf("%d@email.com", i),
			PasswordHash: helpers.CreatePasswordHash(fmt.Sprintf("%d", i)),
			UserGroupID:  models.USER,
			Status:       models.ACTIVE,
		}
	}
	userActivationToken = helpers.CreateSecureToken()
	usersDBMock.Users = append(usersDBMock.Users, models.User{
		ID:           111,
		Name:         "inactive",
		Email:        "inactive@email.com",
		PasswordHash: helpers.CreatePasswordHash("inactive"),
		UserGroupID:  models.USER,
		Status:       models.NOTACTIVE,
		Token:        userActivationToken,
	})
	usersDBMock.Users = append(usersDBMock.Users, models.User{
		ID:           333,
		Name:         "Admin",
		Email:        "admin@email.com",
		PasswordHash: helpers.CreatePasswordHash("admin"),
		UserGroupID:  models.ADMIN,
		Status:       models.ACTIVE,
	})
}

func seedCategoriesDB() {
	categoriesDBMock.Categories = make([]models.Category, 10)
	for i := 0; i < 10; i++ {
		categoriesDBMock.Categories[i] = models.Category{
			ID:     uint64(i),
			Name:   fmt.Sprintf("User%d", i),
			UserID: authenticatedUser.ID,
		}
	}
}

func seedProjectsDB() {
	projectsDBMock.Projects = make([]models.Project, 10)
	for i := 0; i < 10; i++ {
		projectsDBMock.Projects[i] = models.Project{
			ID:     uint64(i),
			Name:   fmt.Sprintf("Project %d", i),
			UserID: authenticatedUser.ID,
		}
	}
}

func seedCommentsDB() {
	commentsDBMock.Comments = make([]models.Comment, 10)
	for i := 0; i < 10; i++ {
		commentsDBMock.Comments[i] = models.Comment{
			ID:       uint64(i),
			Contents: fmt.Sprintf("Contents %d", i),
			TaskID:   uint64(i),
			UserID:   authenticatedUser.ID,
		}
	}
}

func seedPagesDB() {
	pagesDBMock.Pages = make([]models.Page, 10)
	for i := 0; i < 10; i++ {
		pagesDBMock.Pages[i] = models.Page{
			ID:        uint64(i),
			Name:      fmt.Sprintf("Page %d", i),
			Published: true,
		}
	}
}

func seedSessionsDB() {
	sessionsDBMock.Sessions = make([]models.Session, 10)
	for i := 0; i < 10; i++ {
		sessionsDBMock.Sessions[i] = models.Session{
			ID:       uint64(i),
			Contents: fmt.Sprintf("Session %d", i),
			UserID:   authenticatedUser.ID,
		}
	}
}

func seedTasksDB() {
	tasksDBMock.Tasks = make([]models.Task, 10)
	for i := 0; i < 10; i++ {
		tasksDBMock.Tasks[i] = models.Task{
			ID:     uint64(i),
			Name:   fmt.Sprintf("Task%d", i),
			UserID: authenticatedUser.ID,
		}
	}
}

func jsonPost(url string, body interface{}) (*http.Response, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(j))
	return resp, err
}

func jsonPut(url string, body interface{}) (*http.Response, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest("PUT", url, bytes.NewReader(j))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	return client.Do(request)
}

func jsonDelete(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	return client.Do(request)
}

func jsonGet(url string) (*http.Response, error) {
	return http.Get(url)
}

func jsonGetAuth(url string, jwt string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
	return client.Do(request)
}

func jsonPostAuth(url string, body interface{}, jwt string) (*http.Response, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(j))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
	return client.Do(request)
}

func jsonPutAuth(url string, body interface{}, jwt string) (*http.Response, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest("PUT", url, bytes.NewReader(j))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
	return client.Do(request)
}

func jsonDeleteAuth(url string, jwt string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
	return client.Do(request)
}

func getAuthenticatedUser() models.User {
	u, err := usersDBMock.GetByEmail("9@email.com")
	if err != nil {
		panic(err)
	}
	if err := u.CreateJWTToken(); err != nil {
		panic(err)
	}
	return u
}

func getAuthenticatedAdmin() models.User {
	u, err := usersDBMock.GetByEmail("admin@email.com")
	if err != nil {
		panic(err)
	}
	if err := u.CreateJWTToken(); err != nil {
		panic(err)
	}
	return u
}
