package seeder

import (
	"TixTrain/app/model"
	"TixTrain/database"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
)

func SeedSchedules() error {
	var scheduleGroups []model.ScheduleGroup
	if err := database.DB.Find(&scheduleGroups).Error; err != nil {
		return err
	}

	var stations []model.Station
	if err := database.DB.Where("latitude IS NOT NULL AND longitude IS NOT NULL").Find(&stations).Error; err != nil {
		return err
	}

	if len(scheduleGroups) == 0 || len(stations) == 0 {
		log.Println("No schedule groups or stations found")
		return nil
	}

	log.Printf("Seeding schedules for %d schedule groups...", len(scheduleGroups))

	batchSize := 500
	schedules := make([]model.Schedule, 0, batchSize)

	for _, sg := range scheduleGroups {
		// Setiap schedule group punya 3-8 stasiun
		numStationsSlice, err := faker.RandomInt(3, 8)
		if err != nil {
			return err
		}
		numStations := numStationsSlice[0]

		// Pilih stasiun random
		selectedStations := make([]model.Station, 0, numStations)
		usedIndices := make(map[int]bool)

		for len(selectedStations) < numStations {
			idxSlice, err := faker.RandomInt(0, len(stations)-1)
			if err != nil {
				return err
			}
			idx := idxSlice[0]
			if !usedIndices[idx] {
				selectedStations = append(selectedStations, stations[idx])
				usedIndices[idx] = true
			}
		}

		// Buat jadwal untuk 90 hari ke depan (setiap hari ada 1 keberangkatan)
		for day := 0; day < 90; day++ {
			baseDate := time.Now().AddDate(0, 0, day)

			// Waktu keberangkatan random antara jam 05:00 - 20:00
			departureHourSlice, err := faker.RandomInt(5, 20)
			if err != nil {
				return err
			}
			departureHour := departureHourSlice[0]

			departureMinuteSlice, err := faker.RandomInt(0, 59)
			if err != nil {
				return err
			}
			departureMinute := departureMinuteSlice[0]

			currentTime := time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), departureHour, departureMinute, 0, 0, time.Local)

			for order, station := range selectedStations {
				var arrivalTime, departureTime time.Time

				if order == 0 {
					// Stasiun pertama: hanya ada departure
					arrivalTime = currentTime
					departureTime = currentTime
				} else if order == len(selectedStations)-1 {
					// Stasiun terakhir: hanya ada arrival
					travelTimeSlice, err := faker.RandomInt(30, 180)
					if err != nil {
						return err
					}
					travelTime := travelTimeSlice[0]
					currentTime = currentTime.Add(time.Duration(travelTime) * time.Minute)
					arrivalTime = currentTime
					departureTime = currentTime
				} else {
					// Stasiun tengah: ada arrival dan departure
					travelTimeSlice, err := faker.RandomInt(30, 180)
					if err != nil {
						return err
					}
					travelTime := travelTimeSlice[0]
					currentTime = currentTime.Add(time.Duration(travelTime) * time.Minute)
					arrivalTime = currentTime

					stopTimeSlice, err := faker.RandomInt(5, 15)
					if err != nil {
						return err
					}
					stopTime := stopTimeSlice[0]
					departureTime = currentTime.Add(time.Duration(stopTime) * time.Minute)
					currentTime = departureTime
				}

				schedule := model.Schedule{
					ScheduleGroupID: sg.ID,
					StationID:       station.ID,
					ArrivalTime:     arrivalTime,
					DepartureTime:   departureTime,
					Order:           order + 1,
					CreatedAt:       time.Now().AddDate(0, 0, -1),
					UpdatedAt:       time.Now().AddDate(0, 0, -1),
				}

				schedules = append(schedules, schedule)

				// Batch insert
				if len(schedules) >= batchSize {
					if err := database.DB.CreateInBatches(schedules, batchSize).Error; err != nil {
						return err
					}
					log.Printf("Inserted %d schedules", len(schedules))
					schedules = schedules[:0]
				}
			}
		}
	}

	// Insert remaining
	if len(schedules) > 0 {
		if err := database.DB.CreateInBatches(schedules, batchSize).Error; err != nil {
			return err
		}
		log.Printf("Inserted final %d schedules", len(schedules))
	}

	log.Println("Schedules seeding completed")
	return nil
}
