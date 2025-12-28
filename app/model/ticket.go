package model

import "time"

type Ticket struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	ScheduleID uint
	Schedule   Schedule `gorm:"foreignKey:ScheduleID"`

	SeatID uint
	Seat   Seat `gorm:"foreignKey:SeatID"`

	Price  int64
	Status string

	CreatedAt time.Time
	UpdatedAt time.Time
}
