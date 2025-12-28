package model

import "time"

type Station struct {
	ID        uint `Gorm:"primaryKey"`
	Name      string
	Code      string `Gorm:"unique"`
	Longitude float64
	Latitude  float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
