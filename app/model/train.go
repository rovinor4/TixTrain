package model

import "time"

type Train struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null,size:255"`
	Code      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
