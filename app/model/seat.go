package model

import "time"

type Seat struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	CoachID uint
	Coach   Coach `gorm:"foreignKey:CoachID"`

	Number    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
