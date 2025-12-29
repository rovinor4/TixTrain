package model

import "time"

type IdentityCard struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Type   string `json:"type" gorm:"size:50;default:'KTP'"`
	Number string `json:"number" gorm:"size:100;not null" faker:"number:20"`

	Name   string `json:"name" gorm:"size:200;not null" faker:"name"`
	Gender string `json:"gender" gorm:"size:10;not null;default:'Male'" faker:"oneof:'Male','Female'"`

	UserID uint
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
