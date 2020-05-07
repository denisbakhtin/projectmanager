package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//PeriodicitiesDB is a periodicities db repository
var PeriodicitiesDB PeriodicitiesRepository

func init() {
	PeriodicitiesDB = &periodicitiesRepository{}
}

//PeriodicitiesRepository is a repository of periodicities
type PeriodicitiesRepository interface {
	CreateRecurringTasks(date time.Time) error
}

type periodicitiesRepository struct{}

//GetAll returns settings for all periodical tasks for all users
func (pr *periodicitiesRepository) CreateRecurringTasks(date time.Time) error {
	var periodicities []Periodicity
	query := db.Where("periodicity_type > 0")
	query = query.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Order("tasks.completed desc, tasks.id desc")
	})
	err := query.Find(&periodicities).Error
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		for _, per := range periodicities {
			//if no tasks or there is atleast one open task, then skip, to prevent spamming
			if len(per.Tasks) == 0 || per.Tasks[len(per.Tasks)-1].Completed == false {
				continue
			}

			task := per.Tasks[0]
			//create a fresh task if today is The day

			switch per.PeriodicityType {
			case DAILY:
				newTask := freshTask(task, date, date)
				if err := tx.Create(&newTask).Error; err != nil {
					return err
				}
			case WEEKLY:
				//bitwise check
				if (weekdayMask(time.Now().Weekday()) & per.Weekdays) > 0 {
					newTask := freshTask(task, date, date)
					if err := tx.Create(&newTask).Error; err != nil {
						return err
					}
				}
			case MONTHLY:
				if date.Day() == per.RepeatingFrom.Day() {
					endDate := time.Date(date.Year(), date.Month(), per.RepeatingTo.Day(), 0, 0, 0, 0, nil)
					newTask := freshTask(task, date, endDate)
					if err := tx.Create(&newTask).Error; err != nil {
						return err
					}
				}
			case YEARLY:
				if date.Day() == per.RepeatingFrom.Day() && date.Month() == per.RepeatingFrom.Month() {
					endDate := time.Date(date.Year(), per.RepeatingTo.Month(), per.RepeatingTo.Day(), 0, 0, 0, 0, nil)
					newTask := freshTask(task, date, endDate)
					if err := tx.Create(&newTask).Error; err != nil {
						return err
					}
				}
			default:
			}

		}
		return nil
	})
	return err
}

func freshTask(task Task, startDate time.Time, endDate time.Time) Task {
	newTask := task
	newTask.ID = 0
	newTask.StartDate = &startDate
	newTask.EndDate = &endDate
	newTask.Completed = false
	newTask.PeriodicityID = task.PeriodicityID
	return newTask
}

func weekdayMask(day time.Weekday) uint {
	//in Go sunday == 0, but my weeks start at Monday
	if day == time.Sunday {
		day = 7
	}
	return 2 << (day - 1)
}
