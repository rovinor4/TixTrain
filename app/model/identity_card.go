package model

import "time"

type IdentityCard struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Type        string     `json:"type" gorm:"size:50;default:'KTP'"`
	Number      string     `json:"number" gorm:"size:100;not null" faker:"number:20"`
	Name        string     `json:"name" gorm:"size:200;not null" faker:"name"`
	Gender      string     `json:"gender" gorm:"size:10;not null;default:'Male'" faker:"oneof:'Male','Female'"`
	DateOfBirth *time.Time `json:"date_of_birth" gorm:"nullable"`
	IsMe        bool       `json:"is_me" gorm:"default:false;index"`
	UserID      uint       `json:"user_id" gorm:"index;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
