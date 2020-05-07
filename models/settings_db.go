package models

//SettingsDB is a settings db repository
var SettingsDB SettingsRepository

func init() {
	SettingsDB = &settingsRepository{}
}

//SettingsRepository is a repository of settings
type SettingsRepository interface {
	GetAll() ([]Setting, error)
	Get(id interface{}) (Setting, error)
	Create(setting Setting) (Setting, error)
	Update(setting Setting) (Setting, error)
	Delete(id interface{}) error
}

type settingsRepository struct{}

//GetAll returns all settings owned by specified user
func (cr *settingsRepository) GetAll() ([]Setting, error) {
	var settings []Setting
	err := db.Order("id").Find(&settings).Error
	return settings, err
}

//Get fetches a setting by its id
func (cr *settingsRepository) Get(id interface{}) (Setting, error) {
	setting := Setting{}
	err := db.First(&setting, id).Error
	return setting, err
}

//Create creates a new setting in db
func (cr *settingsRepository) Create(setting Setting) (Setting, error) {
	err := db.Create(&setting).Error
	return setting, err
}

//Update updates a setting in db
func (cr *settingsRepository) Update(setting Setting) (Setting, error) {
	err := db.Save(&setting).Error
	return setting, err
}

//Delete removes a setting from db
func (cr *settingsRepository) Delete(id interface{}) error {
	setting := Setting{}
	err := db.First(&setting, id).Error
	if err != nil {
		return err
	}
	return db.Delete(&setting).Error
}
