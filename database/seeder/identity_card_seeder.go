package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
)

func SeedIdentityCards() error {
	// Check if identity cards already exist
	var count int64
	if err := pkg.DB.Model(&model.IdentityCard{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Printf("Identity cards already exist (%d records). Skipping seeder.", count)
		return nil
	}

	// Fetch all users with role passenger
	var users []model.User
	if err := pkg.DB.Select("id").Where("role", "passenger").Find(&users).Error; err != nil {
		return err
	}

	if len(users) == 0 {
		log.Println("No users found for identity cards")
		return nil
	}

	batchSize := 500
	identityTypes := []string{"KTP", "SIM", "Passport"}

	log.Printf("Seeding identity cards for %d passenger users...", len(users))

	// Each passenger will have 1-3 identity cards
	identityCards := make([]model.IdentityCard, 0, len(users)*2)
	genders := []string{"Male", "Female"}

	for idx, user := range users {
		// Random number of identity cards (1-3)
		numCards := rand.Intn(3) + 1

		for i := 0; i < numCards; i++ {
			identityType := identityTypes[rand.Intn(len(identityTypes))]

			// Generate realistic number based on type
			var number string

			switch identityType {
			case "KTP":
				provinsi := rand.Intn(34) + 1 // 1-34 provinsi di Indonesia
				kabupaten := rand.Intn(99) + 1
				kecamatan := rand.Intn(99) + 1
				tanggal := rand.Intn(31) + 1
				bulan := rand.Intn(12) + 1
				tahun := rand.Intn(50) + 50 // 1950-1999
				urut := rand.Intn(9999) + 1
				number = fmt.Sprintf("%02d%02d%02d%02d%02d%02d%04d",
					provinsi, kabupaten, kecamatan, tanggal, bulan, tahun, urut)
			case "SIM":
				tahunTerbit := rand.Intn(10) + 14 // 2014-2023
				nomorUrut := rand.Intn(99999999) + 1
				number = fmt.Sprintf("%02d%010d", tahunTerbit, nomorUrut)
			case "Passport":
				huruf := string(rune('A' + rand.Intn(26)))
				angka := rand.Intn(9999999) + 1
				number = fmt.Sprintf("%s%07d", huruf, angka)
			}

			// Generate random name and gender
			name := faker.Name()
			gender := genders[rand.Intn(len(genders))]

			daysAgo := rand.Intn(365) + 1
			createdAt := time.Now().AddDate(0, 0, -daysAgo)
			updatedAt := createdAt.AddDate(0, 0, rand.Intn(daysAgo))

			identityCard := model.IdentityCard{
				Type:      identityType,
				Number:    number,
				Name:      name,
				Gender:    gender,
				UserID:    user.ID,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			}

			identityCards = append(identityCards, identityCard)
		}

		// Batch insert
		if (idx+1)%batchSize == 0 || idx == len(users)-1 {
			if err := pkg.DB.Create(&identityCards).Error; err != nil {
				return err
			}
			log.Printf("Progress: %d/%d users processed", idx+1, len(users))
			identityCards = make([]model.IdentityCard, 0, len(users)*2)
		}
	}

	log.Println("Identity cards seeding completed")
	return nil
}
