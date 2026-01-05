package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type UintArray []uint

func (a UintArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *UintArray) Scan(value interface{}) error {
	if value == nil {
		*a = UintArray{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, a)
}

type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = StringArray{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, a)
}

type ScheduleGroup struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `json:"name"`
	TrainID uint
	Train   Train `gorm:"foreignKey:TrainID" json:"train"`

	DepartureStations UintArray   `gorm:"type:jsonb" json:"departure_stations"`
	ArrivalStations   UintArray   `gorm:"type:jsonb" json:"arrival_stations"`
	Class             StringArray `gorm:"type:jsonb" json:"class"` // ekonomi, bisnis, eksekutif

	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`

	Schedules []Schedule `gorm:"foreignKey:ScheduleGroupID" json:"schedules"`

	Coaches []Coach `gorm:"foreignKey:ScheduleGroupID" json:"coaches"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
