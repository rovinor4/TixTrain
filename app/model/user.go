package model

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type User struct {
	ID              uint             `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string           `json:"name" faker:"name"`
	Email           string           `json:"email" gorm:"unique;size:255;not null" faker:"email,unique"`
	Password        string           `json:"password" gorm:"size:255;not null"`
	ProfilePicture  *string          `json:"profile_picture" gorm:"size:255;nullable"`
	VectorPicture   *pgvector.Vector `json:"vector_picture" gorm:"type:vector(1536);nullable"`
	Role            string           `json:"role" gorm:"size:50;default:'passenger'"` // roles: passenger, admin,conductor
	EmailVerifiedAt *time.Time       `json:"email_verified_at" gorm:"nullable"`
	IdentityCard    []IdentityCard   `gorm:"foreignKey:UserID;"`
	StationId       *uint            `json:"station_id" gorm:"nullable"`
	Stations        Station          `gorm:"foreignKey:StationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
