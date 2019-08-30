package models

import "time"

//Notification represents a user notification, after it's being read it is deleted
type Notification struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	UserID    uint64    `json:"user_id"`
	Entity    string    `json:"entity"`
	EntityID  uint64    `json:"entity_id"`
}
