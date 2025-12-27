package model

import "time"

type Seat struct {
	ID        uint `Gorm:"primaryKey"`
	CoachID   uint
	Number    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
