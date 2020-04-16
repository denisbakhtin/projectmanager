package controllers

import (
	"html/template"
	"log"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//ListenAndServe prepares routes & runs web server
func ListenAndServe(mode string) {

	//++++++++++++++++++++ GIN ROUTES +++++++++++++++++++++++++++++
	gin.SetMode(mode)

	router := setupRouter()

	log.Printf("Starting application in %s mode", mode)
	router.Run(":8181")
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Static("/public", "./public")
	router.SetFuncMap(funcMap())
	router.LoadHTMLGlob("views/**/*")
	router.Use(LogErrors())
	router.GET("/", home)
	router.GET("/pages/:id", pagesGetHTML)

	publicAPI := router.Group("/api")
	{
		publicAPI.POST("/login", loginPost)
		publicAPI.POST("/activate", activatePost)
		publicAPI.POST("/register", registerPost)
		publicAPI.POST("/forgot", forgotPost)
		publicAPI.POST("/reset", resetPost)
		publicAPI.GET("/settings", settingsGet)
	}

	api := router.Group("/api")
	api.Use(AuthRequired())
	{
		/*
			api.GET("/roles", rolesGet)
			api.GET("/roles/:id", roleGet)
			api.POST("/roles", rolesPost)
			api.PUT("/roles/:id", rolesPut)
			api.DELETE("/roles/:id", rolesDelete)
		*/

		api.GET("/account", accountGet)
		api.PUT("/account", accountPut)

		api.GET("/reports/spent", spentGet)

		api.GET("/categories", categoriesGet)
		api.GET("/categories/:id", categoryGet)
		api.POST("/categories", categoriesPost)
		api.PUT("/categories/:id", categoriesPut)
		api.DELETE("/categories/:id", categoriesDelete)
		api.GET("/categories_summary", categoriesSummaryGet)

		api.GET("/sessions", sessionsGet)
		api.GET("/sessions/:id", sessionGet)
		api.GET("/new_session", sessionNewGet)
		api.POST("/sessions", sessionsPost)
		api.DELETE("/sessions/:id", sessionsDelete)
		api.GET("/sessions_summary", sessionsSummaryGet)

		api.POST("/task_logs", taskLogsPost)
		api.PUT("/task_logs/:id", taskLogsPut)

		api.GET("/projects", projectsGet)
		api.GET("/projects/:id", projectGet)
		api.GET("/new_project", projectNewGet)
		api.GET("/edit_project/:id", projectEditGet)
		api.POST("/projects", projectsPost)
		api.PUT("/projects/:id", projectsPut)
		api.DELETE("/projects/:id", projectsDelete)
		api.GET("/favorite_projects", projectsFavoriteGet)
		api.PUT("/archive_project/:id", projectArchive)
		api.PUT("/favor_project/:id", projectFavorite)
		api.GET("/projects_summary", projectsSummaryGet)

		api.GET("/tasks", tasksGet)
		api.GET("/tasks/:id", taskGet)
		api.GET("/new_task", taskNewGet)
		api.GET("/edit_task/:id", taskEditGet)
		api.POST("/tasks", tasksPost)
		api.PUT("/tasks/:id", tasksPut)
		api.DELETE("/tasks/:id", tasksDelete)
		api.GET("/tasks_summary", tasksSummaryGet)

		api.GET("/comments", commentsGet)
		api.GET("/comments/:id", commentGet)
		api.POST("/comments", commentsPost)
		api.PUT("/comments/:id", commentsPut)
		api.DELETE("/comments/:id", commentsDelete)

		/*
			api.GET("/project_users/:project_id", projectUsersGet)
			api.GET("/project_users/:project_id/:id", projectUserGet)
			api.POST("/project_users/:project_id", projectUsersPost)
			api.PUT("/project_users/:project_id/:id", projectUsersPut)
			api.DELETE("/project_users/:project_id/:id", projectUsersDelete)
		*/

		api.POST("/upload/:uploader", uploadsPost)

		api.GET("/search", searchGet)
	}

	//access for admins only
	api.Use(AdminRequired())
	{
		api.GET("/users", usersGet)
		api.GET("/users/:id", userGet)
		api.PUT("/users/:id", userStatusPut)
		api.GET("/users_summary", usersSummaryGet)

		api.GET("/pages", pagesGet)
		api.GET("/pages/:id", pageGet)
		api.POST("/pages", pagesPost)
		api.PUT("/pages/:id", pagesPut)
		api.DELETE("/pages/:id", pagesDelete)

		api.GET("/settings/:id", settingGet)
		api.POST("/settings", settingsPost)
		api.PUT("/settings/:id", settingsPut)
		api.DELETE("/settings/:id", settingsDelete)
	}
	return router
}

//currentUserID returns authenticated user ID
func currentUserID(c *gin.Context) uint64 {
	if u, exists := c.Get("user"); exists {
		if user, ok := u.(models.User); ok {
			return user.ID
		}
	}
	return 0
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"pages":    pagesMenu,
		"siteName": siteName,
	}
}

func pagesMenu() []models.Page {
	pages, _ := models.PagesDB.GetPagesForMenu()
	return pages
}

func siteName() string {
	return config.Settings.ProjectName
}

func abortWithError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, err.Error())
	c.Error(err)
}
