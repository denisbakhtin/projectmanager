package models

import "time"

//AttachedFile represents a polymorphic file attachment
type AttachedFile struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	OwnerID   uint64     `json:"owner_id" gorm:"index:owner" valid:"required"`
	OwnerType string     `json:"owner_type" gorm:"index:owner" valid:"required"`
	Name      string     `json:"name"`
	FilePath  string     `json:"file_path" valid:"required"`
	URL       string     `json:"url" valid:"required"`
}
