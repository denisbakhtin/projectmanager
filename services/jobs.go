package services

import (
	"log"
	"time"

	"github.com/denisbakhtin/projectmanager/models"
	"github.com/jinzhu/gorm"
)

//CreatePeriodicTasks creates new periodic tasks for the specified day.
//Atm only if the previous task has been solved the new one is created (to prevent spam).
//But this will be a subject to change via project or user settings
//All requests are made in transaction and any occurring errors should be logged outside and sent to admin
func CreatePeriodicTasks(date time.Time) error {
	var periodicities []models.Periodicity
	count := 0
	query := models.DB.Where("periodicity_type > 0")
	query = query.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Order("tasks.id desc").Limit(1)
	})
	if err := query.Find(&periodicities).Error; err != nil {
		return err
	}
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		for _, per := range periodicities {
			if len(per.Tasks) == 0 || per.Tasks[0].Completed == false {
				continue
			}

			task := per.Tasks[0]
			//create a fresh task if today is The day

			switch per.PeriodicityType {
			case models.DAILY:
				newTask := freshTask(task, date, date)
				if err := tx.Create(&newTask).Error; err != nil {
					return err
				}
				count++
			case models.WEEKLY:
				//bitwise check
				if (weekdayMask(time.Now().Weekday()) & per.Weekdays) > 0 {
					newTask := freshTask(task, date, date)
					if err := tx.Create(&newTask).Error; err != nil {
						return err
					}
					count++
				}
			case models.MONTHLY:
				if date.Day() == per.RepeatingFrom.Day() {
					endDate := time.Date(date.Year(), date.Month(), per.RepeatingTo.Day(), 0, 0, 0, 0, nil)
					newTask := freshTask(task, date, endDate)
					if err := tx.Create(&newTask).Error; err != nil {
						return err
					}
					count++
				}
			case models.YEARLY:
				if date.Day() == per.RepeatingFrom.Day() && date.Month() == per.RepeatingFrom.Month() {
					endDate := time.Date(date.Year(), per.RepeatingTo.Month(), per.RepeatingTo.Day(), 0, 0, 0, 0, nil)
					newTask := freshTask(task, date, endDate)
					if err := tx.Create(&newTask).Error; err != nil {
						return err
					}
					count++
				}
			default:
			}

		}
		return nil
	})
	if err != nil {
		return err
	}

	log.Printf("%d periodic task(s) have been created.\n", count)
	return nil
}

func freshTask(task models.Task, startDate time.Time, endDate time.Time) models.Task {
	newTask := task
	newTask.ID = 0
	newTask.StartDate = startDate
	newTask.EndDate = endDate
	newTask.Completed = false
	return newTask
}

func weekdayMask(day time.Weekday) uint {
	//in Go sunday == 0, but my weeks start at Monday
	if day == time.Sunday {
		day = 7
	}
	return 2 << (day - 1)
}
