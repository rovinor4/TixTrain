package model

import "time"

type RolePermission struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID       uint       `json:"role_id" gorm:"not null;index"`
	PermissionID uint       `json:"permission_id" gorm:"not null;index"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
