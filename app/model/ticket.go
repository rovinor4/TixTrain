package model

import "time"

type Ticket struct {
	ID uint `Gorm:"primaryKey"`

	UserID uint
	User   User `Gorm:"foreignKey:UserID"`

	ScheduleID uint
	Schedule   Schedule `Gorm:"foreignKey:ScheduleID"`

	SeatID uint
	Seat   Seat `Gorm:"foreignKey:SeatID"`

	Price  int64
	Status string

	CreatedAt time.Time
	UpdatedAt time.Time
}
