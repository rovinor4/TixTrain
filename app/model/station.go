package model

import "time"

type Station struct {
	ID        uint `Gorm:"primaryKey"`
	Name      string
	Code      string `Gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
