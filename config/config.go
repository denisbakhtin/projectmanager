package config

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var (
	//Settings keeps db settings for current environment
	Settings settingYaml
	//SettingsAll keeps db settings for all environments
	SettingsAll settingsYaml
	//AppDir is the full path to application dir
	AppDir string
	//AssetsDir is the full path to assets folder
	AssetsDir string
	//UploadPath is the full path to uploads dir
	UploadPath string
	//UploadPathURL is the uploads dir url (/public/uploads)
	UploadPathURL string
	//LogFile is log file reference, opened by config.Initialize, closed by main.main's defer
	LogFile *os.File
)

//settings for ALL environments, see config.yml
type settingsYaml struct {
	Production  settingYaml
	Test        settingYaml
	Development settingYaml
}

//settings for one environment
type settingYaml struct {
	Connection   string //postgresql database connection string
	SMTPServer   string `yaml:"smtp_server"`   //smtp server address
	SMTPLogin    string `yaml:"smtp_login"`    //smtp login
	SMTPPassword string `yaml:"smtp_password"` //smtp user password
	SMTPPort     int    `yaml:"smtp_port"`     //smtp port
	SMTPReply    string `yaml:"smtp_reply"`    //reply to email address
	ProjectName  string `yaml:"project_name"`  //Project name (for page titles, emails subjects and so on)
	JWTSecret    string `yaml:"jwt_secret"`    //Secret string, used to sign jwt tokens
	CheckEmails  bool   `yaml:"check_emails"`  //Check email existence with github.com/badoux/checkmail (for production mainly)
}

//Initialize loads config file and initializes config variables
func Initialize(mode string) {
	AppDir, _ = filepath.Abs("")
	switch mode {
	case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
	default:
		log.Printf("Wrong value of -mode flag: %s, setting it to 'debug'", mode)
		mode = gin.DebugMode
	}

	UploadPathURL = "/public/uploads"
	UploadPath = path.Join(AppDir, "public", "uploads")
	AssetsDir = path.Join(AppDir, "public", "assets")

	var err error
	//closed in main.main by defer
	LogFile, err = os.OpenFile(path.Join(AppDir, "logs", mode+".log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v. All logs will be redirected to STDOUT", err)
	}
	//redirect standard log output to file
	log.SetOutput(LogFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	SettingsAll = loadDbSettings()
	switch mode {
	case gin.DebugMode:
		Settings = SettingsAll.Development
	case gin.ReleaseMode:
		Settings = SettingsAll.Production
	case gin.TestMode:
		Settings = SettingsAll.Test
	}
}

func loadDbSettings() settingsYaml {
	configFile := path.Join(AppDir, "config", "config.yml")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Panic(err.Error())
	}

	buffer, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panic(err.Error())
	}

	allSettings := settingsYaml{}
	err = yaml.Unmarshal(buffer, &allSettings)
	if err != nil {
		log.Panic(err.Error())
	}
	return allSettings
}
