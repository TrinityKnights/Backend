package entity

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" gorm:"not null"`
	Time        string `json:"time" gorm:"not null"`
	VenueID     uint      `json:"venue_id" gorm:"not null"`
	Venue       Venue     `json:"venue" gorm:"foreignKey:VenueID"`
	gorm.Model
}

func (e *Event) TableName() string {
	return "events"
}
