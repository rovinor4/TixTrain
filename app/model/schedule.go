package model

import "time"

type Schedule struct {
	ID              uint `Gorm:"primaryKey"`
	ScheduleGroupID uint
	ScheduleGroup   ScheduleGroup `Gorm:"foreignKey:ScheduleGroupID"`

	StationID uint
	Station   Station `Gorm:"foreignKey:StationID"`

	ArrivalTime   time.Time
	DepartureTime time.Time
	Order         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
