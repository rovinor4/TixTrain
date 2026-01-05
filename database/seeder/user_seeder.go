package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
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

	// Check if users already exist
	var count int64
	if err := pkg.DB.Model(&model.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Printf("Users already exist (%d records). Skipping seeder.", count)
		return nil
	}

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
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Default Staff",
			Email:     "staff@example.com",
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Default Admin",
			Email:     "admin@example.com",
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := pkg.DB.Create(&defaultUsers).Error; err != nil {
		return err
	}

	// Assign roles to default users
	roleMap := map[string]string{
		"passenger@example.com": "passenger",
		"staff@example.com":     "staff",
		"admin@example.com":     "admin",
	}

	for _, user := range defaultUsers {
		if roleName, ok := roleMap[user.Email]; ok {
			var role model.Role
			if err := pkg.DB.Where("name = ?", roleName).First(&role).Error; err == nil {
				userRole := model.UserRole{
					UserID: user.ID,
					RoleID: role.ID,
				}
				pkg.DB.Create(&userRole)
			}
		}
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

			user := model.User{
				Name:            faker.Name(),
				Email:           faker.Email(),
				Password:        string(hashedPassword),
				ProfilePicture:  nil,
				EmailVerifiedAt: emailVerifiedAt,
				CreatedAt:       time.Now().AddDate(0, 0, -createdOffset),
				UpdatedAt:       time.Now().AddDate(0, 0, -updatedOffset),
			}

			users = append(users, user)
		}

		if err := pkg.DB.CreateInBatches(users, batchSize).Error; err != nil {
			return err
		}

		// Assign roles to created users
		for idx, user := range users {
			var roleName string
			if batch*batchSize+idx < 200 {
				roleName = "staff"
			} else if batch*batchSize+idx < 300 {
				roleName = "admin"
			} else {
				roleName = "passenger"
			}

			var role model.Role
			if err := pkg.DB.Where("name = ?", roleName).First(&role).Error; err == nil {
				userRole := model.UserRole{
					UserID: user.ID,
					RoleID: role.ID,
				}
				pkg.DB.Create(&userRole)
			}
		}

		log.Printf("Progress: %d/%d users end", (batch+1)*batchSize, totalUsers)
	}

	log.Println("Users seeding completed")
	return nil
}
