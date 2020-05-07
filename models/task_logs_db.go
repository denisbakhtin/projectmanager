package models

//TaskLogsDB is a task logs db repository
var TaskLogsDB TaskLogsRepository

func init() {
	TaskLogsDB = &taskLogsRepository{}
}

//TaskLogsRepository is a taskLogs repository
type TaskLogsRepository interface {
	Create(userID uint64, taskLog TaskLog) (TaskLog, error)
	Update(userID uint64, taskLog TaskLog) (TaskLog, error)
}

type taskLogsRepository struct{}

//Create preates a new taskLog in db
func (r *taskLogsRepository) Create(userID uint64, taskLog TaskLog) (TaskLog, error) {
	taskLog.UserID = userID
	taskLog.SessionID = 0
	err := db.Create(&taskLog).Error
	return taskLog, err
}

//Update updates a taskLog in db
func (r *taskLogsRepository) Update(userID uint64, taskLog TaskLog) (TaskLog, error) {
	taskLog.UserID = userID
	taskLog.SessionID = 0
	err := db.Save(&taskLog).Error
	return taskLog, err
}
