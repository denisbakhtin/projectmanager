package models

//ReportsDB is a reports db repository
var ReportsDB ReportsRepository

func init() {
	ReportsDB = &reportsRepository{}
}

//ReportsRepository is a reports repository
type ReportsRepository interface {
	Spent(userID uint64) ([]TaskLog, error)
}

type reportsRepository struct{}

//Spent returns spent report data
func (r *reportsRepository) Spent(userID uint64) ([]TaskLog, error) {
	var taskLogs []TaskLog

	err := DB.Preload("Task.Project").Preload("Task").
		Preload("Task.Category").
		Where("session_id = 0 and minutes > 0 and user_id = ?", userID).
		Find(&taskLogs).Error

	return taskLogs, err
}
