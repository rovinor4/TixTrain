package model

import "time"

type Permission struct {
	ID        uint             `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string           `json:"name" gorm:"unique;size:50;not null"`
	Route     string           `json:"route" gorm:"size:255;not null"`
	Roles     []RolePermission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
