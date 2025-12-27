package model

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type User struct {
	ID              uint `Gorm:"primaryKey"`
	Name            string
	Email           string `Gorm:"unique"`
	EmailVerifiedAt *time.Time
	Password        string
	RememberToken   *string
	profilePicture  string
	VectorPicture   pgvector.Vector `Gorm:"type:vector(1536);column:vector_picture"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
