package mocks

import "github.com/denisbakhtin/projectmanager/models"

type ReportsDBMock struct {
	Logs []models.TaskLog
}

//Spent returns spent report data
func (r *ReportsDBMock) Spent(userID uint64) ([]models.TaskLog, error) {
	return r.Logs, nil
}
