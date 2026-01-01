package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"fmt"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
)

func SeedScheduleGroupsAndCoaches() error {
	var trains []model.Train
	if err := pkg.DB.Find(&trains).Error; err != nil {
		return err
	}

	if len(trains) == 0 {
		return fmt.Errorf("no trains found, please run train seeder first")
	}

	log.Printf("Seeding schedule groups and coaches for %d trains...", len(trains))

	coachClasses := []string{"Ekonomi", "Bisnis", "Eksekutif"}

	for _, train := range trains {
		// Buat 2-5 schedule group per train (rute berbeda)
		numScheduleGroups, err := faker.RandomInt(2, 5)
		if err != nil {
			return err
		}

		for sg := 0; sg < numScheduleGroups[0]; sg++ {
			createdAtOffset, err := faker.RandomInt(30, 365)
			if err != nil {
				return err
			}
			updatedAtOffset, err := faker.RandomInt(0, 30)
			if err != nil {
				return err
			}

			scheduleGroup := model.ScheduleGroup{
				Name:      fmt.Sprintf("%s - Route %d", train.Name, sg+1),
				TrainID:   train.ID,
				CreatedAt: time.Now().AddDate(0, 0, -createdAtOffset[0]),
				UpdatedAt: time.Now().AddDate(0, 0, -updatedAtOffset[0]),
			}

			if err := pkg.DB.Create(&scheduleGroup).Error; err != nil {
				return err
			}
		}

		// Buat 5-10 gerbong per train
		numCoaches, err := faker.RandomInt(5, 10)
		if err != nil {
			return err
		}

		for c := 0; c < numCoaches[0]; c++ {
			createdAtOffset, err := faker.RandomInt(30, 365)
			if err != nil {
				return err
			}
			updatedAtOffset, err := faker.RandomInt(0, 30)
			if err != nil {
				return err
			}

			coach := model.Coach{
				TrainID:   train.ID,
				Code:      fmt.Sprintf("%s-%d", train.Code, c+1),
				Class:     coachClasses[c%len(coachClasses)],
				CreatedAt: time.Now().AddDate(0, 0, -createdAtOffset[0]),
				UpdatedAt: time.Now().AddDate(0, 0, -updatedAtOffset[0]),
			}

			if err := pkg.DB.Create(&coach).Error; err != nil {
				return err
			}
		}
	}

	log.Println("Schedule groups and coaches seeding completed")
	return nil
}
