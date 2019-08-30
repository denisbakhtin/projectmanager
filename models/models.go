package models

import (
	"fmt"
	"log"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"

	//required by GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DB references sqlx db connection
var DB *gorm.DB

//InitializeDB initializes DB variable by establishing connection to the DB server
func InitializeDB() {
	var err error
	DB, err = gorm.Open("postgres", config.Settings.Connection)
	if err != nil {
		log.Panic(err)
	}
	validations.RegisterCallbacks(DB)
	DB.AutoMigrate(&User{}, &UserGroup{}, &Task{}, &Project{}, &Status{}, &TaskLog{}, &TaskStep{},
		&AttachedFile{}, &ProjectUser{}, &Role{}, &Page{}, &Log{}, &Notification{})
	DB.AutoMigrate()

	count := 0
	if err := DB.Model(&UserGroup{}).Where([]int64{ADMIN, EDITOR, USER}).Count(&count).Error; err != nil {
		log.Panic(err)
	}
	if count != 3 {
		if err := DB.Model(&UserGroup{}).Create(&UserGroup{ID: ADMIN, Name: "Admin"}).Error; err != nil {
			log.Panic(fmt.Sprintf("Error creating Admin user group with ID=%d. Try to create it manually and modify models.ADMIN constant accordingly.", ADMIN))
		}
		if err := DB.Model(&UserGroup{}).Create(&UserGroup{ID: EDITOR, Name: "Editor"}).Error; err != nil {
			log.Panic(fmt.Sprintf("Error creating Editor user group with ID=%d. Try to create it manually and modify models.EDITOR constant accordingly.", EDITOR))
		}
		if err := DB.Model(&UserGroup{}).Create(&UserGroup{ID: USER, Name: "User"}).Error; err != nil {
			log.Panic(fmt.Sprintf("Error creating User user group with ID=%d. Try to create it manually and modify models.USER constant accordingly.", USER))
		}
	}
}
