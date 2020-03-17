package models

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/fiam/gounidecode/unidecode"
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
	DB.AutoMigrate(&User{}, &UserGroup{}, &Task{}, &Project{}, &TaskLog{},
		&AttachedFile{}, &Page{}, &Log{}, &Notification{},
		&Setting{}, &Category{}, &Session{}, &Comment{}, &TaskLog{})

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

//createSlug makes url slug out of string
func createSlug(s string) string {
	s = strings.ToLower(unidecode.Unidecode(s))                     //transliterate if it is not in english
	s = regexp.MustCompile("[^a-z0-9\\s]+").ReplaceAllString(s, "") //spaces
	s = regexp.MustCompile("\\s+").ReplaceAllString(s, "-")         //spaces
	return s
}
