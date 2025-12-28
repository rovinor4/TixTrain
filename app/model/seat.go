package model

import "time"

type Seat struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CoachID   uint
	Number    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
