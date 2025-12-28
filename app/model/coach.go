package model

import "time"

type Coach struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	TrainID uint
	Train   Train `gorm:"foreignKey:TrainID"`

	Code  string
	Class string

	Seats []Seat `gorm:"foreignKey:CoachID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
