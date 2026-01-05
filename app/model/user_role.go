package model

import "time"

type UserRole struct {
	ID        uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint `json:"user_id" gorm:"not null;index"`
	RoleID    uint `json:"role_id" gorm:"not null;index"`
	User      User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role      Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
