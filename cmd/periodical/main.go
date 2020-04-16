package main

import (
	"log"
	"time"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/denisbakhtin/projectmanager/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Initialize(gin.ReleaseMode) //settings
	if config.LogFile != nil {
		defer config.LogFile.Close()
	}
	models.InitializeDB()
	defer models.Close()

	if err := services.CreatePeriodicTasks(time.Now()); err != nil {
		log.Printf("Error executing CreatePeriodicTasks: %v\n", err)
	}
}
