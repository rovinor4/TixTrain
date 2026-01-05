package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"fmt"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
)

func SeedCoaches() error {
	var scheduleGroups []model.ScheduleGroup
	if err := pkg.DB.Preload("Train").Find(&scheduleGroups).Error; err != nil {
		return err
	}

	if len(scheduleGroups) == 0 {
		return fmt.Errorf("no schedule groups found, please run schedule seeder first")
	}

	log.Printf("Seeding coaches for %d schedule groups...", len(scheduleGroups))

	batchSize := 100
	coaches := make([]model.Coach, 0, batchSize)

	for _, scheduleGroup := range scheduleGroups {
		// Validasi schedule group memiliki class
		if len(scheduleGroup.Class) == 0 {
			log.Printf("Warning: schedule group %d has no classes, skipping", scheduleGroup.ID)
			continue
		}

		// Buat coaches berdasarkan class yang tersedia di schedule group
		classes := scheduleGroup.Class

		// Buat 2-4 gerbong per class
		for _, className := range classes {
			if className == "" {
				log.Printf("Warning: empty class name in schedule group %d, skipping", scheduleGroup.ID)
				continue
			}

			numCoachesPerClass, err := faker.RandomInt(2, 4)
			if err != nil {
				return err
			}

			for c := 0; c < numCoachesPerClass[0]; c++ {
				createdAtOffset, err := faker.RandomInt(30, 365)
				if err != nil {
					return err
				}
				updatedAtOffset, err := faker.RandomInt(0, 30)
				if err != nil {
					return err
				}

				price := getCoachPrice(className)

				coach := model.Coach{
					ScheduleGroupID: scheduleGroup.ID,
					Code:            fmt.Sprintf("%s-%s-%d", scheduleGroup.Train.Code, className[:3], c+1),
					Class:           className,
					Price:           price,
					CreatedAt:       time.Now().AddDate(0, 0, -createdAtOffset[0]),
					UpdatedAt:       time.Now().AddDate(0, 0, -updatedAtOffset[0]),
				}

				coaches = append(coaches, coach)

				// Batch insert
				if len(coaches) >= batchSize {
					if err := pkg.DB.Create(&coaches).Error; err != nil {
						return err
					}
					log.Printf("Inserted %d coaches", len(coaches))
					coaches = coaches[:0]
				}
			}
		}
	}

	// Insert remaining coaches
	if len(coaches) > 0 {
		if err := pkg.DB.Create(&coaches).Error; err != nil {
			return err
		}
		log.Printf("Inserted final %d coaches", len(coaches))
	}

	log.Println("Coaches seeding completed")
	return nil
}

func getCoachPrice(class string) int64 {
	prices := map[string]int64{
		"Ekonomi":   50000,
		"Bisnis":    100000,
		"Eksekutif": 150000,
	}

	if price, exists := prices[class]; exists {
		return price
	}
	return 50000
}
