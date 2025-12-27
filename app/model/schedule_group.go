package model

import "time"

type ScheduleGroup struct {
	ID      uint `Gorm:"primaryKey"`
	Name    string
	TrainID uint
	Train   Train `Gorm:"foreignKey:TrainID"`

	Schedules []Schedule `Gorm:"foreignKey:ScheduleGroupID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
