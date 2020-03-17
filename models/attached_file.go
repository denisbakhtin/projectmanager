package models

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/jinzhu/gorm"
)

//AttachedFile represents a polymorphic file attachment
type AttachedFile struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	OwnerID   uint64    `json:"owner_id" gorm:"index:owner" valid:"required"`
	OwnerType string    `json:"owner_type" gorm:"index:owner" valid:"required"`
	UserID    uint64    `json:"user_id" valid:"-"`
	User      User      `json:"user" gorm:"save_associations:false" valid:"-"`
	Name      string    `json:"name"`
	FilePath  string    `json:"file_path" valid:"required"`
	URL       string    `json:"url" valid:"required"`
}

// BeforeDelete - gorm hook, validate removal here & delete other related records (no cascade tag atm)
func (a *AttachedFile) BeforeDelete(tx *gorm.DB) (err error) {
	if err = validateFilePath(a.FilePath); err != nil {
		return
	}
	if err = os.Remove(a.FilePath); err != nil {
		return
	}
	return
}

//validateFilePath checks if attachment file path is valid
func validateFilePath(filePath string) error {
	//check it's not empty and does not lead outside of project folder
	if len(filePath) == 0 || len(path.Ext(filePath)) == 0 || strings.Contains(filePath, "..") || strings.Contains(filePath, "./") || strings.ContainsAny(filePath, "*?") {
		return fmt.Errorf("Attached file has incorrect path")
	}
	//check its inside public/uploads subdirectory
	if !strings.HasPrefix(filePath, config.UploadPath) {
		return fmt.Errorf("File does not belong the correct uploads directory")
	}
	return nil
}
