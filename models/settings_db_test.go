package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestSettingsRepositoryGetAll(t *testing.T) {
	list, err := SettingsDB.GetAll()
	assert.Nil(t, err)
	assert.Greater(t, len(list), 0)
}

func TestSettingsRepositoryGet(t *testing.T) {
	s := getOrCreateSetting()
	setting, err := SettingsDB.Get(s.ID)
	assert.Nil(t, err)
	assert.NotEmpty(t, setting)
	assert.NotZero(t, setting.ID)
}

func TestSettingsRepositoryCreate(t *testing.T) {
	code := fmt.Sprintf("setting_%d", time.Now().Nanosecond())
	setting := Setting{
		Code:  code,
		Value: code,
	}
	all, _ := SettingsDB.GetAll()
	count := len(all)
	setting, err := SettingsDB.Create(setting)
	assert.Nil(t, err)
	assert.NotZero(t, setting.ID)
	assert.Equal(t, setting.Code, code)

	setting, err = SettingsDB.Get(setting.ID)
	assert.Nil(t, err)
	assert.NotZero(t, setting.ID)
	assert.Equal(t, setting.Code, code)

	all, _ = SettingsDB.GetAll()
	assert.Equal(t, count+1, len(all), "The number of settings should have been increased by 1")
}

func TestSettingsRepositoryUpdate(t *testing.T) {
	code := fmt.Sprintf("setting_%d", time.Now().Nanosecond())
	setting := Setting{}
	DB.First(&setting)
	assert.NotZero(t, setting.ID)
	assert.NotEqual(t, setting.Code, code)

	setting.Code = code
	p, err := SettingsDB.Update(setting)
	assert.Nil(t, err)
	assert.Equal(t, p.Code, code)
}

func TestSettingsRepositoryDelete(t *testing.T) {
	err := SettingsDB.Delete(11111111111)
	assert.NotNil(t, err)
	assert.True(t, gorm.IsRecordNotFoundError(err), "This setting should not exist")

	setting := getOrCreateSetting()
	assert.NotZero(t, setting.ID)

	err = SettingsDB.Delete(setting.ID)
	assert.Nil(t, err)
	_, err = SettingsDB.Get(setting.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Setting should have been removed")
}
