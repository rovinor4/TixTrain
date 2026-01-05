package model

import "time"

type Seat struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	CoachID uint
	Coach   Coach `gorm:"foreignKey:CoachID" json:"coach"`

	Number    string    `json:"number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
