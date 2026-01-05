package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"log"
	"time"
)

func SeedPermissions() error {
	log.Println("Seeding permissions...")

	permissions := []model.Permission{
		// Admin permissions
		{
			Name:      "view_admin_stats",
			Route:     "GET /admin/stats",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "view_admin_tickets",
			Route:     "GET /admin/tickets",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "view_admin_tickets_all",
			Route:     "GET /admin/tickets/all",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "manage_admin_stations",
			Route:     "GET /admin/stations",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		// Staff permissions
		{
			Name:      "view_staff_tickets",
			Route:     "GET /staff/tickets",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		// Passenger permissions
		{
			Name:      "view_passenger_tickets",
			Route:     "GET /passenger/tickets",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Check if permissions already exist
	var count int64
	if err := pkg.DB.Model(&model.Permission{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("Permissions already exist, skipping seeding")
		return nil
	}

	if err := pkg.DB.Create(&permissions).Error; err != nil {
		return err
	}

	log.Println("Permissions seeding completed")
	return nil
}
