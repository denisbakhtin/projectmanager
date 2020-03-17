package main

import (
	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/controllers"
	"github.com/denisbakhtin/projectmanager/models"
)

func main() {
	config.Initialize() //settings
	if config.LogFile != nil {
		defer config.LogFile.Close()
	}
	models.InitializeDB()
	defer models.DB.Close()

	controllers.ListenAndServe()
}
