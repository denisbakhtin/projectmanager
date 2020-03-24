package models

import (
	"errors"
	"time"

	"github.com/denisbakhtin/projectmanager/helpers"
	"github.com/jinzhu/gorm"
)

//Task periodicity type
const (
	DONTREPEAT = iota // == 0
	DAILY             //== 1
	WEEKLY            // == 2, etc
	MONTHLY
	YEARLY
)

//Periodicity represents a periodicity settings for periodical tasks
type Periodicity struct {
	ID              uint64    `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PeriodicityType uint      `json:"periodicity_type,string"` //see constants
	Weekdays        uint      `json:"weekdays"`                //a bitmask of selected weekdays, makes sense only for periodicity == WEEKLY. 1 - Monday ... 64 - Sunday
	RepeatingFrom   time.Time `json:"repeating_from"`          //for periodicity == MONTHLY or YEARLY. This will be the StartDate and EndDate (for each month or year) of the newly created task accordingly
	RepeatingTo     time.Time `json:"repeating_to"`            //see RepeatingFrom
	Tasks           []Task    `json:"tasks" gorm:"save_associations:false" valid:"-"`
	UserID          uint64    `json:"user_id" valid:"-"`
	User            User      `json:"user" gorm:"save_associations:false" valid:"-"`
}

// BeforeUpdate - gorm hook, fired before record update
func (p *Periodicity) BeforeUpdate(tx *gorm.DB) (err error) {
	//check if original user_id is not being changed
	per := Periodicity{}
	if tx.Where("id = ? and user_id = ?", p.ID, p.UserID).First(&per); per.ID == 0 {
		return errors.New(helpers.NotFoundOrOwned("Periodicity settings"))
	}
	return
}
