package model

import "time"

type TicketDetail struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	TicketID uint
	Ticket   Ticket `gorm:"foreignKey:TicketID"`

	SeatID uint
	Seat   Seat `gorm:"foreignKey:SeatID"`

	IdentityCardID uint
	IdentityCard   IdentityCard `gorm:"foreignKey:IdentityCardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
