package controllers

import (
	"html/template"
	"path"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/gin-gonic/gin"
)

var (
	tpl *template.Template //shared inside controllers package
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
	router.GET("/", homeHandler)
	router.GET("/admin", adminHandler)
	router.GET("/pages/:id", pagesGetHTMLHandler)
	router.POST("/api/login", loginPostHandler)
	router.POST("/api/activate", activatePostHandler)
	router.POST("/api/register", registerPostHandler)
	router.POST("/api/forgot", forgotPostHandler)
	router.POST("/api/reset", resetPostHandler)

	authorized := router.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/api/roles", rolesGetHandler)
		authorized.GET("/api/roles/:id", roleGetHandler)
		authorized.POST("/api/roles", rolesPostHandler)
		authorized.PUT("/api/roles/:id", rolesPutHandler)
		authorized.DELETE("/api/roles/:id", rolesDeleteHandler)
		authorized.GET("/api/users", usersGetHandler)
		authorized.GET("/api/users/:id", userGetHandler)
		authorized.PUT("/api/users/:id", usersPutHandler)
		authorized.GET("/api/account", accountGetHandler)
		authorized.PUT("/api/account", accountPutHandler)

		authorized.GET("/api/pages", pagesGetHandler)
		authorized.GET("/api/pages/:id", pageGetHandler)
		authorized.POST("/api/pages", pagesPostHandler)
		authorized.PUT("/api/pages/:id", pagesPutHandler)

		authorized.GET("/api/projects", projectsGetHandler)
		authorized.GET("/api/projects/:id", projectGetHandler)
		authorized.POST("/api/projects", projectsPostHandler)
		authorized.PUT("/api/projects/:id", projectsPutHandler)
		authorized.DELETE("/api/projects/:id", projectsDeleteHandler)

		authorized.GET("/api/statuses", statusesGetHandler)
		authorized.GET("/api/statuses/:id", statusGetHandler)
		authorized.POST("/api/statuses", statusesPostHandler)
		authorized.PUT("/api/statuses/:id", statusesPutHandler)
		authorized.DELETE("/api/statuses/:id", statusesDeleteHandler)

		authorized.GET("/api/task_steps", taskStepsGetHandler)
		authorized.GET("/api/task_steps/:id", taskStepGetHandler)
		authorized.POST("/api/task_steps", taskStepsPostHandler)
		authorized.PUT("/api/task_steps/:id", taskStepsPutHandler)
		authorized.DELETE("/api/task_steps/:id", taskStepsDeleteHandler)

		authorized.GET("/api/tasks", tasksGetHandler)
		authorized.GET("/api/tasks/:id", taskGetHandler)
		authorized.POST("/api/tasks", tasksPostHandler)
		authorized.PUT("/api/tasks/:id", tasksPutHandler)
		authorized.DELETE("/api/tasks/:id", tasksDeleteHandler)

		authorized.GET("/api/project_users/:project_id", projectUsersGetHandler)
		authorized.GET("/api/project_users/:project_id/:id", projectUserGetHandler)
		authorized.POST("/api/project_users/:project_id", projectUsersPostHandler)
		authorized.PUT("/api/project_users/:project_id/:id", projectUsersPutHandler)
		authorized.DELETE("/api/project_users/:project_id/:id", projectUsersDeleteHandler)

		authorized.POST("/api/upload/:uploader", uploadsPostHandler)
	}

	router.Run(":44555")

	//load template files all at once
	tpl = template.Must(template.ParseGlob(path.Join(config.AppDir, "views", "*", "*.tmpl")))
}
