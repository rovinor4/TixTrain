package model

import "time"

type Coach struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	ScheduleGroupID uint
	ScheduleGroup   ScheduleGroup `gorm:"foreignKey:ScheduleGroupID"`

	Code  string
	Class string
	Price int64

	Seats []Seat `gorm:"foreignKey:CoachID"`
	Quota int    `gorm:"default:100"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
