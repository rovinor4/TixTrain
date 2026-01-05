package model

import "time"

type Role struct {
	ID          uint             `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string           `json:"name" gorm:"unique;size:50;not null"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Users       []User           `gorm:"many2many:user_roles;foreignKey:ID;joinForeignKey:RoleID;References:ID;joinReferences:UserID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
