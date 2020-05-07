package models

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/denisbakhtin/marble/models"
	"github.com/denisbakhtin/projectmanager/config"
	"github.com/fiam/gounidecode/unidecode"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"

	//required by GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//InitializeDB initializes db variable by establishing connection to the db server
func InitializeDB() {
	var err error
	db, err = gorm.Open("postgres", config.Settings.Connection)
	if err != nil {
		log.Panic(err)
	}
	validations.RegisterCallbacks(db)
	db.AutoMigrate(&User{}, &UserGroup{}, &Task{}, &Project{}, &TaskLog{},
		&AttachedFile{}, &Page{}, &Log{},
		&Setting{}, &Category{}, &Session{}, &Comment{}, &TaskLog{}, &Periodicity{})

	runManualMigrations()

	count := 0
	if err := db.Model(&UserGroup{}).Where([]int64{ADMIN, EDITOR, USER}).Count(&count).Error; err != nil {
		log.Panic(err)
	}
	if count != 3 {
		if err := db.Create(&UserGroup{ID: ADMIN, Name: "Admin"}).Error; err != nil {
			log.Panic(fmt.Sprintf("Error creating Admin user group with ID=%d. Try to create it manually and modify models.ADMIN constant accordingly.", ADMIN))
		}
		if err := db.Create(&UserGroup{ID: EDITOR, Name: "Editor"}).Error; err != nil {
			log.Panic(fmt.Sprintf("Error creating Editor user group with ID=%d. Try to create it manually and modify models.EDITOR constant accordingly.", EDITOR))
		}
		if err := db.Create(&UserGroup{ID: USER, Name: "User"}).Error; err != nil {
			log.Panic(fmt.Sprintf("Error creating User user group with ID=%d. Try to create it manually and modify models.USER constant accordingly.", USER))
		}
	}
	//ensure site_name setting exists in database
	setting := models.Setting{}
	if err := db.Where("code = ?", "site_name").First(&setting).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			if err := db.Create(&models.Setting{Code: "site_name", Value: config.Settings.ProjectName}).Error; err != nil {
				log.Panic(fmt.Sprintf("Can't create site_name setting in database. %v", err))
			}
		} else {
			log.Panic(err)
		}
	}
}

//Close terminates db handler
func Close() {
	db.Close()
}

func runManualMigrations() {
	if err := db.Exec("update tasks set start_date = null where start_date::date = '0001-01-01'").Error; err != nil {
		log.Panic(err)
	}
	if err := db.Exec("update tasks set end_date = null where end_date::date = '0001-01-01'").Error; err != nil {
		log.Panic(err)
	}
	if err := db.Exec("update periodicities set repeating_from = null where repeating_from::date = '0001-01-01'").Error; err != nil {
		log.Panic(err)
	}
	if err := db.Exec("update periodicities set repeating_to = null where repeating_to::date = '0001-01-01'").Error; err != nil {
		log.Panic(err)
	}
}

//createSlug makes url slug out of string
func createSlug(s string) string {
	s = strings.ToLower(unidecode.Unidecode(s))                     //transliterate if it is not in english
	s = regexp.MustCompile("[^a-z0-9\\s]+").ReplaceAllString(s, "") //spaces
	s = regexp.MustCompile("\\s+").ReplaceAllString(s, "-")         //spaces
	return s
}
