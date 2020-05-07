package services

import (
	"time"

	"github.com/denisbakhtin/projectmanager/models"
)

//CreatePeriodicTasks creates new periodic tasks for the specified day.
//Atm only if the previous task has been solved the new one is created (to prevent spam).
//But this will be a subject to change via project or user settings
//All requests are made in transaction and any occurring errors should be logged outside and sent to admin
func CreatePeriodicTasks(date time.Time) error {
	return models.PeriodicitiesDB.CreateRecurringTasks(date)
}
