package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/controllers"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

func main() {
	mode := flag.String("mode", "debug", fmt.Sprintf("application mode, one of: %s, %s, %s", gin.DebugMode, gin.ReleaseMode, gin.TestMode))
	flag.Parse()
	if *mode != gin.DebugMode && *mode != gin.TestMode && *mode != gin.ReleaseMode {
		log.Panicf("Wrong application mode %s\n", *mode)
	}

	config.Initialize(*mode) //settings
	defer config.LogFile.Close()

	models.InitializeDB()
	defer models.Close()

	controllers.ListenAndServe(*mode)
}
