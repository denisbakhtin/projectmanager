package models

import (
	"crypto/sha1"
	"fmt"
	"time"
)

//ParticipantRequest represents a row from participant_requests table
type ParticipantRequest struct {
	IDFrom    uint64    `json:"id_from" valid:"required"`
	EmailTo   string    `json:"email_to" valid:"required,email"`
	Token     string    `json:"token" valid:"required,length(1|500)"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

//BeforeCreate gorm hook
func (r *ParticipantRequest) BeforeCreate() (err error) {
	tmpStr := fmt.Sprintf("%d-%v-%v", r.IDFrom, r.EmailTo, time.Now().UnixNano())
	r.Token = fmt.Sprintf("%x%x%x", r.IDFrom, r.EmailTo, sha1.Sum([]byte(tmpStr)))
	return
}

/*
func GetParticipantRequestByIds(id_from, email_to interface{}) *ParticipantRequest {
	record := &ParticipantRequest{}
	DB.Get(record, "SELECT * FROM participant_requests WHERE id_from=$1 AND email_to=$2", id_from, email_to)
	return record
}

//GetParticipantRequestListByIdFrom returns a list of participant requests, given id_from
func GetParticipantRequestListByIdFrom(id interface{}) ([]ParticipantRequest, error) {
	list := make([]ParticipantRequest, 0)
	var err error
	err = DB.Select(&list, "SELECT * FROM participant_requests WHERE id_from=$1 ORDER BY created_at ASC", id)
	return list, err
}

//GetParticipantRequestListByEmailTo returns a list of participant requests, given email_to
func GetParticipantRequestListByEmailTo(email string) ([]ParticipantRequest, error) {
	list := make([]ParticipantRequest, 0)
	var err error
	err = DB.Select(&list, "SELECT * FROM participant_requests WHERE email_to=$1 ORDER BY created_at ASC", email)
	return list, err
}

//GetParticipantRequestUserListByIdFrom returns a list of users related to participant id
func GetParticipantRequestUserListByIdFrom(id interface{}) ([]User, error) {
	list := make([]User, 0)
	var err error
	err = DB.Select(&list, "SELECT id, name, email, status FROM users WHERE email IN (SELECT email_to FROM participant_requests WHERE id_from=$1) ORDER BY id ASC", id)
	return list, err
}

//GetParticipantRequestUserListByEmailTo returns a list of users related to participant id
func GetParticipantRequestUserListByEmailTo(email string) ([]User, error) {
	list := make([]User, 0)
	var err error
	err = DB.Select(&list, "SELECT id, name, email, status FROM users WHERE id IN (SELECT id_from FROM participant_requests WHERE email_to=$1) ORDER BY id ASC", email)
	return list, err
}
*/
