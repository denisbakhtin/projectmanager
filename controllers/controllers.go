package controllers

import (
	"github.com/denisbakhtin/projectmanager/config"
	"github.com/gin-gonic/gin"
)

//Initialize prepares handlers, loads template files into memory, inits sessionStore
func Initialize() {

	//++++++++++++++++++++ GIN ROUTES +++++++++++++++++++++++++++++
	router := gin.Default() //with logger and recover middlewares baked in
	if config.Env != "production" {
		//in production assets are served by nginx
		router.Static("/public", "./public")
	}
	router.LoadHTMLGlob("views/**/*")
	router.GET("/", home)
	router.GET("/pages/:id", pagesGetHTML)
	router.POST("/api/login", loginPost)
	router.POST("/api/activate", activatePost)
	router.POST("/api/register", registerPost)
	router.POST("/api/forgot", forgotPost)
	router.POST("/api/reset", resetPost)

	authorized := router.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/api/roles", rolesGet)
		authorized.GET("/api/roles/:id", roleGet)
		authorized.POST("/api/roles", rolesPost)
		authorized.PUT("/api/roles/:id", rolesPut)
		authorized.DELETE("/api/roles/:id", rolesDelete)

		authorized.GET("/api/users", usersGet)
		authorized.GET("/api/users/:id", userGet)
		authorized.PUT("/api/users/:id", usersPut)
		authorized.GET("/api/account", accountGet)
		authorized.PUT("/api/account", accountPut)

		authorized.GET("/api/pages", pagesGet)
		authorized.GET("/api/pages/:id", pageGet)
		authorized.POST("/api/pages", pagesPost)
		authorized.PUT("/api/pages/:id", pagesPut)

		authorized.GET("/api/projects", projectsGet)
		authorized.GET("/api/projects/:id", projectGet)
		authorized.GET("/api/new_project", projectNewGet)
		authorized.GET("/api/edit_project/:id", projectEditGet)
		authorized.POST("/api/projects", projectsPost)
		authorized.PUT("/api/projects/:id", projectsPut)
		authorized.DELETE("/api/projects/:id", projectsDelete)

		authorized.GET("/api/statuses", statusesGet)
		authorized.GET("/api/statuses/:id", statusGet)
		authorized.POST("/api/statuses", statusesPost)
		authorized.PUT("/api/statuses/:id", statusesPut)
		authorized.DELETE("/api/statuses/:id", statusesDelete)

		authorized.GET("/api/task_steps", taskStepsGet)
		authorized.GET("/api/task_steps/:id", taskStepGet)
		authorized.POST("/api/task_steps", taskStepsPost)
		authorized.PUT("/api/task_steps/:id", taskStepsPut)
		authorized.DELETE("/api/task_steps/:id", taskStepsDelete)

		authorized.GET("/api/tasks", tasksGet)
		authorized.GET("/api/tasks/:id", taskGet)
		authorized.POST("/api/tasks", tasksPost)
		authorized.PUT("/api/tasks/:id", tasksPut)
		authorized.DELETE("/api/tasks/:id", tasksDelete)

		authorized.GET("/api/project_users/:project_id", projectUsersGet)
		authorized.GET("/api/project_users/:project_id/:id", projectUserGet)
		authorized.POST("/api/project_users/:project_id", projectUsersPost)
		authorized.PUT("/api/project_users/:project_id/:id", projectUsersPut)
		authorized.DELETE("/api/project_users/:project_id/:id", projectUsersDelete)

		authorized.GET("/api/settings", settingsGet)
		authorized.GET("/api/settings/:id", settingGet)
		authorized.POST("/api/settings", settingsPost)
		authorized.PUT("/api/settings/:id", settingsPut)
		authorized.DELETE("/api/settings/:id", settingsDelete)

		authorized.POST("/api/upload/:uploader", uploadsPost)

		authorized.GET("/api/notifications", notificationsGet)
		authorized.DELETE("/api/notifications/:id", notificationsDelete)
	}

	router.Run(":44555")
}
