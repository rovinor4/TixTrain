package model

import "time"

type Ticket struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	ScheduleID uint
	Schedule   Schedule `gorm:"foreignKey:ScheduleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	TicketDetail []TicketDetail `gorm:"foreignKey:TicketID"`

	Price  int64
	Status string

	CreatedAt time.Time
	UpdatedAt time.Time
}
