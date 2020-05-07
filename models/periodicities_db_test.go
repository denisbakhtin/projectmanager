package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPeriodicitiesRepositoryCreateRecurringTasks(t *testing.T) {
	per := createPeriodicity()
	var tasks []Task
	err := db.Where("periodicity_id = ?", per.ID).Find(&tasks).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tasks))
	err = PeriodicitiesDB.CreateRecurringTasks(time.Now())
	assert.Nil(t, err)
	err = db.Where("periodicity_id = ?", per.ID).Find(&tasks).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(tasks))
}
