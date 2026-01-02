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
	// select id and name, and correct Where syntax
	if err := pkg.DB.Select("id, name").Where("role = ?", "passenger").Find(&users).Error; err != nil {
		return err
	}

	if len(users) == 0 {
		log.Println("No users found for identity cards")
		return nil
	}

	// seed randomness
	rand.Seed(time.Now().UnixNano())

	identityTypes := []string{"KTP", "SIM", "Passport"}

	log.Printf("Seeding identity cards for %d passenger users...", len(users))

	genders := []string{"Male", "Female"}

	for idx, user := range users {
		// Random number of identity cards (1-3)
		numCards := rand.Intn(3) + 1
		userCards := make([]model.IdentityCard, 0, numCards)

		// Pick which card will be the primary (and ensure its Name == user.Name)
		primaryIdx := rand.Intn(numCards)

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

			// Use user's name for the chosen primary identity, otherwise random name
			var name string
			if i == primaryIdx {
				name = user.Name
			} else {
				name = faker.Name()
			}

			gender := genders[rand.Intn(len(genders))]

			// Generate random date of birth (age 17-80 years)
			yearsOld := rand.Intn(64) + 17 // 17-80 tahun
			daysOffset := rand.Intn(365)
			dateOfBirth := time.Now().AddDate(-yearsOld, 0, -daysOffset)

			daysAgo := rand.Intn(365) + 1
			createdAt := time.Now().AddDate(0, 0, -daysAgo)
			updatedAt := createdAt.AddDate(0, 0, rand.Intn(daysAgo))

			identityCard := model.IdentityCard{
				Type:        identityType,
				Number:      number,
				Name:        name,
				Gender:      gender,
				DateOfBirth: &dateOfBirth,
				IsMe:        i == primaryIdx,
				UserID:      user.ID,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			}

			userCards = append(userCards, identityCard)
		}

		// Insert this user's identity cards
		if err := pkg.DB.Create(&userCards).Error; err != nil {
			return err
		}

		// Progress log per user
		if (idx+1)%100 == 0 || idx == len(users)-1 {
			log.Printf("Progress: %d/%d users processed", idx+1, len(users))
		}
	}

	log.Println("Identity cards seeding completed")
	return nil
}
