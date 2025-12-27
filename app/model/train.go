package model

import "time"

type Train struct {
	ID        uint `Gorm:"primaryKey"`
	Name      string
	Code      string `Gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
