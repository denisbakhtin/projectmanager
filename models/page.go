package models

import (
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/kennygrant/sanitize"
)

//Page represents a record in pages table
type Page struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"-"`
	Name        string     `json:"name" valid:"required, length(1|500)"`
	Description string     `json:"description" valid:"length(0,100000)"`
	Published   bool       `json:"published"`
}

//Excerpt returns page excerpt
func (p *Page) Excerpt(length int) string {
	s := sanitize.HTML(p.Description)
	ss := helpers.Substr(s, 0, length)
	return ss
}
