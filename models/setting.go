package models

import "time"

//Setting type contains settings info
type Setting struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Code  string `json:"code" valid:"required,length(1|100)"`
	Title string `json:"title"`
	Value string `json:"value" valid:"required"`
}
