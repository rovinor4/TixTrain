package seeder

import (
	"TixTrain/app/model"
	"TixTrain/database"
	"log"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() error {
	totalUsers := 10000
	batchSize := 500

	log.Printf("Seeding %d users...", totalUsers)

	// Hash password once
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create default accounts
	defaultUsers := []model.User{
		{
			Name:      "Default Passenger",
			Email:     "passenger@example.com",
			Password:  string(hashedPassword),
			Role:      "passenger",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Default Staff",
			Email:     "staff@example.com",
			Password:  string(hashedPassword),
			Role:      "staff",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Default Admin",
			Email:     "admin@example.com",
			Password:  string(hashedPassword),
			Role:      "admin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := database.DB.Create(&defaultUsers).Error; err != nil {
		return err
	}

	// Generate remaining users
	for batch := 0; batch < totalUsers/batchSize; batch++ {
		users := make([]model.User, 0, batchSize)

		log.Printf("Progress: %d/%d users start", (batch+1)*batchSize, totalUsers)

		for i := 0; i < batchSize; i++ {
			// Random email verified (70% verified)
			var emailVerifiedAt *time.Time
			if i%10 < 7 {
				verifiedOffset := rand.Intn(365) + 1
				verifiedTime := time.Now().AddDate(0, 0, -verifiedOffset)
				emailVerifiedAt = &verifiedTime
			}

			createdOffset := rand.Intn(730) + 1
			updatedOffset := rand.Intn(30) + 1

			// Assign roles based on distribution
			var role string
			if batch*batchSize+i < 200 {
				role = "staff"
			} else if batch*batchSize+i < 300 {
				role = "admin"
			} else {
				role = "passenger"
			}

			user := model.User{
				Name:            faker.Name(),
				Email:           faker.Email(),
				Password:        string(hashedPassword),
				ProfilePicture:  nil, // Profile picture set to nil
				Role:            role,
				EmailVerifiedAt: emailVerifiedAt,
				CreatedAt:       time.Now().AddDate(0, 0, -createdOffset),
				UpdatedAt:       time.Now().AddDate(0, 0, -updatedOffset),
			}

			users = append(users, user)
		}

		if err := database.DB.CreateInBatches(users, batchSize).Error; err != nil {
			return err
		}

		log.Printf("Progress: %d/%d users end", (batch+1)*batchSize, totalUsers)
	}

	log.Println("Users seeding completed")
	return nil
}
