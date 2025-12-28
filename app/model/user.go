package model

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type User struct {
	ID               uint             `gorm:"primaryKey;autoIncrement"`
	Name             string           `faker:"name"`
	Email            string           `gorm:"unique;size:255;not null" faker:"email,unique"`
	Password         string           `gorm:"size:255;not null"`
	ProfilePicture   *string          `gorm:"size:255;nullable"`
	VectorPicture    *pgvector.Vector `gorm:"type:vector(1536);nullable"`
	Role             string           `gorm:"size:50;type:enum('admin','passenger','staff');default:'passenger'"`
	IdentityCardType string           `gorm:"size:50;type:enum('KTP','SIM','Passport');default:'KTP'"`
	IdentityCardNo   string           `gorm:"size:100;not null;unique" faker:"number:20"`
	EmailVerifiedAt  *time.Time       `gorm:"nullable"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
