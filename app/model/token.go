package model

type Token struct {
	ID        uint   `Gorm:"primaryKey"`
	Value     string `Gorm:"unique"`
	UserID    uint
	User      User `Gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	UserAgent string
	CreatedAt int64
	ExpiresAt int64
}
