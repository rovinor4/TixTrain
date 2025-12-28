package model

import "time"

type ScheduleGroup struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Name    string
	TrainID uint
	Train   Train `gorm:"foreignKey:TrainID"`

	Schedules []Schedule `gorm:"foreignKey:ScheduleGroupID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
