package model

import "time"

type Schedule struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	ScheduleGroupID uint
	ScheduleGroup   ScheduleGroup `gorm:"foreignKey:ScheduleGroupID"`

	StationID uint
	Station   Station `gorm:"foreignKey:StationID"`

	ArrivalTime   time.Time
	DepartureTime time.Time
	Order         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
