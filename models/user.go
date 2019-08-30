package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/helpers"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//User status codes
const (
	NOTACTIVE = 0
	ACTIVE    = 1
	SUSPENDED = 2
)

//User represents a row from users table
type User struct {
	ID           uint64     `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"-"`
	Name         string     `json:"name" valid:"required,length(1|100)"`
	Email        string     `json:"email" gorm:"unique_index" valid:"required,email,length(1|100)"`
	PasswordHash string     `json:"-" valid:"required"`
	Token        string     `json:"token" valid:"length(0|1500)"`
	UserGroupID  uint64     `json:"user_group_id" gorm:"index" valid:"required"`
	Status       uint       `json:"status"` //See constants
	UserGroup    UserGroup  `json:"user_group" gorm:"save_associations:false" valid:"-"`
}

//BeforeCreate gorm hook
func (u *User) BeforeCreate() (err error) {
	u.Email = helpers.NormalizeEmail(u.Email)
	return
}

//BeforeUpdate gorm hook
func (u *User) BeforeUpdate() (err error) {
	u.Email = helpers.NormalizeEmail(u.Email)
	return
}

//BeforeDelete gorm hook
func (u *User) BeforeDelete() (err error) {
	if u.IsAdmin() {
		count := 0
		DB.Model(&User{}).Where("user_group_id = ?", ADMIN).Count(&count)
		if count == 1 {
			err = errors.New("Can't remove the last admin")
		}
	}
	return
}

//HasPassword checks if user has this password
func (u *User) HasPassword(password string) bool {
	return helpers.CheckPasswordHash(password, u.PasswordHash)
}

//IsAdmin checks if user is admin
func (u *User) IsAdmin() bool {
	return u.UserGroupID == ADMIN
}

//IsEditor checks if user is editor
func (u *User) IsEditor() bool {
	return u.UserGroupID == EDITOR
}

//IsUser checks if user is a regular user
func (u *User) IsUser() bool {
	return u.UserGroupID == USER
}

//IsActive checks if user is active
func (u *User) IsActive() bool {
	return u.Status == ACTIVE
}

//BelongsToProjectUsers checks if user is among project users
func (u *User) BelongsToProjectUsers(pusers []ProjectUser) bool {
	for i := 0; i < len(pusers); i++ {
		if pusers[i].UserID == u.ID {
			return true
		}
	}
	return false
}

//JWTClaims extends jwt-go.StandardClaims with custom fields
type JWTClaims struct {
	jwt.StandardClaims
	Name   string `json:"name,omitempty"`
	Role   string `json:"role,omitempty"`
	UserID uint64 `json:"user_id,omitempty"`
}

//CreateJWTToken issues a valid JWT token
func (u *User) CreateJWTToken() error {
	DB.First(&u.UserGroup, u.UserGroupID)
	now := time.Now()
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			Id:        createJWTID(u.ID),
			Issuer:    "pm",
			ExpiresAt: now.AddDate(0, 0, 30).Unix(),
			Subject:   u.Email,
		},
		UserID: u.ID,
		Name:   u.Name,
		Role:   u.UserGroup.Name,
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	u.Token, err = rawToken.SignedString([]byte(config.Settings.JWTSecret))
	return err
}

//createJWTID creates a secure jwt token id
func createJWTID(id uint64) string {
	str := fmt.Sprintf("%v-%v", id, time.Now().UnixNano())
	bytes, _ := bcrypt.GenerateFromPassword([]byte(str), 12)
	return string(bytes)
}
