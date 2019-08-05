package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	//Env represents application mode (ex: development, test, production)
	Env string
	//Settings keeps db settings for current environment
	Settings settingYaml
	//SettingsAll keeps db settings for all environments
	SettingsAll settingsYaml
	//AppDir is the full path to application dir
	AppDir string
	//UploadPath is the full path to uploads dir
	UploadPath string
	//UploadPathURL is the uploads dir url (/public/uploads)
	UploadPathURL string
	//LogFile is log file reference, opened by config.Initialize, closed by main.main's defer
	LogFile *os.File
	//Domain name for session option
	Domain string
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
func Initialize() {
	envPtr := flag.String("e", "development", "application environment, one of: development, production, test")
	flag.Parse()

	AppDir, _ = filepath.Abs("")
	switch *envPtr {
	case "development":
		Env = *envPtr
		Domain = "" //empty is ok for dev
	case "test":
		Env = *envPtr
		Domain = ""
	case "production":
		Env = *envPtr
		log.Panic("AppDir & Domain not specified")
	default:
		log.Printf("Wrong value of -e flag: %s, setting it to 'development'", *envPtr)
		Env = "development"
		Domain = ""
	}

	UploadPathURL = "/public/uploads"
	UploadPath = path.Join(AppDir, "public", "uploads")

	//log to file only in production, in dev mode I like to see messages on screen
	if Env == "production" {
		var err error
		//closed in main.main by defer
		LogFile, err = os.OpenFile(path.Join(AppDir, "logs", Env+".log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening log file: %v", err)
		}
		//works around all packages, importing standard "log"!!!! awesome tbh
		log.SetOutput(LogFile)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	DbSettingsAll := loadDbSettings()
	switch Env {
	case "development":
		Settings = DbSettingsAll.Development
	case "production":
		Settings = DbSettingsAll.Production
	case "test":
		Settings = DbSettingsAll.Test
	}

	log.Printf("Starting application in %s mode", Env)
}

func loadDbSettings() *settingsYaml {
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
	return &allSettings
}
