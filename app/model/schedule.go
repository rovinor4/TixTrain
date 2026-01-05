package model

import "time"

type Schedule struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	ScheduleGroupID uint
	ScheduleGroup   ScheduleGroup `gorm:"foreignKey:ScheduleGroupID"`

	DepartureStationID uint
	DepartureStation   Station `gorm:"foreignKey:DepartureStationID"`

	ArrivalStationID uint
	ArrivalStation   Station `gorm:"foreignKey:ArrivalStationID"`

	ArrivalTime   time.Time
	DepartureTime time.Time

	Order     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
