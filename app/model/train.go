package model

import "time"

type Train struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"not null,size:255" json:"name"`
	Code      string `gorm:"unique" json:"code"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
