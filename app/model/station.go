package model

import "time"

type Station struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Code      string `gorm:"unique"`
	Longitude float64
	Latitude  float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
