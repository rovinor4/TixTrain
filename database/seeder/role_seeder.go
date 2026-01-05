package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"log"
	"time"
)

func SeedRoles() error {
	log.Println("Seeding roles...")

	roles := []model.Role{
		{
			Name:      "admin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "staff",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "passenger",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Check if roles already exist
	var count int64
	if err := pkg.DB.Model(&model.Role{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("Roles already exist, skipping seeding")
		return nil
	}

	if err := pkg.DB.Create(&roles).Error; err != nil {
		return err
	}

	log.Println("Roles seeding completed")
	return nil
}
