package models

//SearchVM is a view model for search queries
type SearchVM struct {
	Categories []Category `json:"categories"`
	Projects   []Project  `json:"projects"`
	Tasks      []Task     `json:"tasks"`
	Comments   []Comment  `json:"comments"`
}
