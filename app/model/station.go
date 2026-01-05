package model

import "time"

type Station struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Code      string    `gorm:"unique" json:"code"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	Image     *string   `json:"image" gorm:"type:text;default:null;nullable"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
