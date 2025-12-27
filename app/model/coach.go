package model

import "time"

type Coach struct {
	ID      uint `Gorm:"primaryKey"`
	TrainID uint
	Train   Train `Gorm:"foreignKey:TrainID"`

	Code  string
	Class string

	Seats []Seat `Gorm:"foreignKey:CoachID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
