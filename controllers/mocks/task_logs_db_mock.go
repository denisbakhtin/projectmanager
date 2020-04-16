package mocks

import (
	"fmt"

	"github.com/denisbakhtin/projectmanager/models"
)

//UsersDBMock is a UsersDB repository mock
type TaskLogsDBMock struct {
	TaskLogs []models.TaskLog
}

func (r *TaskLogsDBMock) Create(userID uint64, taskLog models.TaskLog) (models.TaskLog, error) {
	taskLog.UserID = userID
	taskLog.SessionID = 0
	taskLog.ID = 111
	r.TaskLogs = append(r.TaskLogs, taskLog)
	return taskLog, nil
}

//Update updates a taskLog in db
func (r *TaskLogsDBMock) Update(userID uint64, taskLog models.TaskLog) (models.TaskLog, error) {
	for i := range r.TaskLogs {
		if r.TaskLogs[i].ID == taskLog.ID {
			taskLog.UserID = userID
			taskLog.SessionID = 0
			r.TaskLogs[i] = taskLog
			return r.TaskLogs[i], nil
		}
	}
	return models.TaskLog{}, fmt.Errorf("Task log not found")
}
