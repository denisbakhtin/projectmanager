package models

import (
	"html/template"
	"strings"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/kennygrant/sanitize"
	"gopkg.in/russross/blackfriday.v2"
)

//Page represents a record in pages table
type Page struct {
	ID              uint64    `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Name            string    `json:"name" valid:"required, length(1|500)"`
	Description     string    `json:"description" valid:"length(0|100000)"`
	MetaKeywords    string    `json:"meta_keywords" valid:"length(0|200)"`
	MetaDescription string    `json:"meta_description" valid:"length(0|200)"`
	Slug            string    `json:"slug"`
	Published       bool      `json:"published"`
}

//Excerpt returns page excerpt
func (p *Page) Excerpt(length int) string {
	s := sanitize.HTML(p.Description)
	ss := helpers.Substr(s, 0, length)
	return ss
}

//HTMLContent returns parsed html content
func (p *Page) HTMLContent() template.HTML {
	return template.HTML(blackfriday.Run([]byte(p.Description)))
}

//BeforeSave gorm hook
func (p *Page) BeforeSave() (err error) {
	if strings.TrimSpace(p.Slug) == "" {
		p.Slug = createSlug(p.Name)
	}
	return
}
