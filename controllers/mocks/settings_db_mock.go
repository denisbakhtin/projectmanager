package mocks

import (
	"fmt"
	"strconv"

	"github.com/denisbakhtin/projectmanager/models"
)

type SettingsDBMock struct {
	Settings []models.Setting
}

//GetAll returns all settings owned by specified user
func (sm *SettingsDBMock) GetAll() ([]models.Setting, error) {
	return sm.Settings, nil
}

//Get fetches a setting by its id
func (sm *SettingsDBMock) Get(id interface{}) (models.Setting, error) {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range sm.Settings {
		if sm.Settings[i].ID == idi {
			return sm.Settings[i], nil
		}
	}
	return models.Setting{}, fmt.Errorf("Setting not found")
}

//Create smeates a new setting in db
func (sm *SettingsDBMock) Create(setting models.Setting) (models.Setting, error) {
	sm.Settings = append(sm.Settings, setting)
	return setting, nil
}

//Update updates a setting in db
func (sm *SettingsDBMock) Update(setting models.Setting) (models.Setting, error) {
	for i := range sm.Settings {
		if sm.Settings[i].ID == setting.ID {
			sm.Settings[i] = setting
			return setting, nil
		}
	}
	return models.Setting{}, fmt.Errorf("Setting not found")
}

//Delete removes a setting from db
func (sm *SettingsDBMock) Delete(id interface{}) error {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range sm.Settings {
		if sm.Settings[i].ID == idi {
			sm.Settings[i] = sm.Settings[len(sm.Settings)-1] // Copy last element to index i.
			sm.Settings = sm.Settings[:len(sm.Settings)-1]   // Truncate slice.
			return nil
		}
	}
	return fmt.Errorf("Setting not found")
}
