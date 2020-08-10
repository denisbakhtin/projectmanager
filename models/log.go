package models

import "time"

//Log represent log table
//Work in progress
type Log struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IP        string    `json:"ip"`
	Session   string    `json:"session"`
	URL       string    `json:"url"`
	Headers   string    `json:"headers"`
	Message   string    `json:"message"`
}
