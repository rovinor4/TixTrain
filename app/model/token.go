package model

import "time"

type Token struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Value     string `gorm:"unique,not null,size:512"`
	UserID    uint
	User      User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	UserAgent string `gorm:"size:255"`
	CreatedAt time.Time
	ExpiresAt time.Time
}
