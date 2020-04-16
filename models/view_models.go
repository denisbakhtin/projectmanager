package models

//LoginVM is a login view model
type LoginVM struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ActivateVM is an activation view model
type ActivateVM struct {
	Token string `json:"token" binding:"required"`
}

//RegisterVM is a registration view model
type RegisterVM struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ForgotVM is a view model for forgotten password request
type ForgotVM struct {
	Email string `json:"email" binding:"required"`
}

//ResetVM is a view model for password reset requests
type ResetVM struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//AccountVM is an account view model
type AccountVM struct {
	Name            string `json:"name" binding:"required"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

//UsersSummaryVM is a view model for users statistics
type UsersSummaryVM struct {
	Count int `json:"count"`
}

//EditTaskVM is a view model for a new or an edited task
type EditTaskVM struct {
	Projects   []Project  `json:"projects"`
	Categories []Category `json:"categories"`
	Task       `json:"task"`
}

//TasksSummaryVM is a view model for tasks statistics
type TasksSummaryVM struct {
	Count          int       `json:"count"`
	LatestTasks    []Task    `json:"latest_tasks"`
	LatestTaskLogs []TaskLog `json:"latest_task_logs"`
}

//NewSessionVM is a view model for a new session
type NewSessionVM struct {
	TaskLogs []TaskLog `json:"task_logs"`
}

//SessionsSummaryVM is a view model for sessions statistics
type SessionsSummaryVM struct {
	Count int `json:"count"`
}

//SearchVM is a view model for search queries
type SearchVM struct {
	Categories []Category `json:"categories"`
	Projects   []Project  `json:"projects"`
	Tasks      []Task     `json:"tasks"`
	Comments   []Comment  `json:"comments"`
}

//EditProjectVM is a view model for a new or an edited project
type EditProjectVM struct {
	Project    `json:"project"`
	Categories []Category `json:"categories"`
}

//ProjectsSummaryVM is a view model for projects statistics
type ProjectsSummaryVM struct {
	Count int `json:"count"`
}

//CategoriesSummaryVM is a view model for categories statistics
type CategoriesSummaryVM struct {
	Count int `json:"count"`
}
