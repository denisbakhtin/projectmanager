package models

/*
import (
	"time"
)

//Participant represents a row from participants table
type Participant struct {
	ID1       uint64    `json:"id1" gorm:"unique_index:id1id2" valid:"required"`
	ID2       uint64    `json:"id2" gorm:"unique_index:id1id2" valid:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

//BeforeCreate gorm hook
func (p *Participant) BeforeCreate() (err error) {
	if p.ID1 > p.ID2 {
		p.ID1, p.ID2 = p.ID2, p.ID1
	}
	return
}
*/
