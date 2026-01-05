package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
)

// SeedSchedules - Membuat schedule groups dengan rute tetap menggunakan tanggal hari ini
func SeedSchedules() error {
	var trains []model.Train
	if err := pkg.DB.Find(&trains).Error; err != nil {
		return err
	}

	var stations []model.Station
	if err := pkg.DB.Where("latitude IS NOT NULL AND longitude IS NOT NULL").Limit(50).Find(&stations).Error; err != nil {
		return err
	}

	if len(trains) == 0 || len(stations) < 3 {
		log.Println("Not enough trains or stations found")
		return nil
	}

	log.Printf("Seeding schedule groups for %d trains...", len(trains))

	// Gunakan tanggal hari ini
	baseDate := time.Now()

	// Buat beberapa rute populer untuk setiap kereta
	routesPerTrain := 2

	for _, train := range trains {
		for r := 0; r < routesPerTrain; r++ {
			// Pilih 3-6 stasiun untuk rute ini
			numStationsSlice, err := faker.RandomInt(3, 6)
			if err != nil {
				return err
			}
			numStations := numStationsSlice[0]

			// Pilih stasiun random untuk rute
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

			if len(selectedStations) < 3 {
				continue
			}

			// Buat departure dan arrival stations list
			departureStations := make(model.UintArray, 0)
			arrivalStations := make(model.UintArray, 0)

			// Semua stasiun kecuali yang terakhir bisa jadi stasiun keberangkatan
			for i := 0; i < len(selectedStations)-1; i++ {
				departureStations = append(departureStations, selectedStations[i].ID)
			}

			// Semua stasiun kecuali yang pertama bisa jadi stasiun kedatangan
			for i := 1; i < len(selectedStations); i++ {
				arrivalStations = append(arrivalStations, selectedStations[i].ID)
			}

			// Tentukan class yang tersedia (varied - gak harus semua ada)
			// Pilih kombinasi random dari available classes
			availableClasses := []string{"Ekonomi", "Bisnis", "Eksekutif"}
			selectedClasses := getRandomClassCombination(availableClasses)

			// Departure jam 10:00, arrival jam 15:00
			departureDateTime := time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), 10, 0, 0, 0, time.Local)

			// Hitung waktu dengan 30-40 menit per stasiun
			totalStations := len(selectedStations)
			intervalMinutes := 30 + (40-30)/totalStations // Pembagi rata
			totalTravelMinutes := intervalMinutes * (totalStations - 1)

			arrivalDateTime := departureDateTime.Add(time.Duration(totalTravelMinutes) * time.Minute)

			// Buat schedule group untuk hari ini
			scheduleGroup := model.ScheduleGroup{
				Name:              train.Name + " - " + selectedStations[0].Name + " to " + selectedStations[len(selectedStations)-1].Name,
				TrainID:           train.ID,
				DepartureStations: departureStations,
				ArrivalStations:   arrivalStations,
				Class:             selectedClasses,
				DepartureTime:     departureDateTime,
				ArrivalTime:       arrivalDateTime,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			if err := pkg.DB.Create(&scheduleGroup).Error; err != nil {
				log.Printf("Error creating schedule group: %v", err)
				continue
			}

			// Buat detail schedule untuk setiap stasiun
			schedules := make([]model.Schedule, 0)
			currentTime := departureDateTime
			intervalMinutesInt := 30 + (40-30)/totalStations

			for order, station := range selectedStations {
				var arrivalTime, departTime time.Time
				var depStationID, arrStationID uint

				if order == 0 {
					// Stasiun pertama
					arrivalTime = currentTime
					departTime = currentTime
					depStationID = station.ID
					if order+1 < len(selectedStations) {
						arrStationID = selectedStations[order+1].ID
					}
				} else if order == len(selectedStations)-1 {
					// Stasiun terakhir
					currentTime = currentTime.Add(time.Duration(intervalMinutesInt) * time.Minute)
					arrivalTime = currentTime
					departTime = currentTime
					depStationID = selectedStations[order-1].ID
					arrStationID = station.ID
				} else {
					// Stasiun tengah
					currentTime = currentTime.Add(time.Duration(intervalMinutesInt) * time.Minute)
					arrivalTime = currentTime
					departTime = currentTime.Add(5 * time.Minute) // Stop 5 menit
					currentTime = departTime

					depStationID = station.ID
					if order+1 < len(selectedStations) {
						arrStationID = selectedStations[order+1].ID
					}
				}

				schedule := model.Schedule{
					ScheduleGroupID:    scheduleGroup.ID,
					DepartureStationID: depStationID,
					ArrivalStationID:   arrStationID,
					ArrivalTime:        arrivalTime,
					DepartureTime:      departTime,
					Order:              order + 1,
					CreatedAt:          time.Now(),
					UpdatedAt:          time.Now(),
				}

				schedules = append(schedules, schedule)
			}

			// Insert semua schedule untuk schedule group ini
			if len(schedules) > 0 {
				if err := pkg.DB.Create(&schedules).Error; err != nil {
					log.Printf("Error creating schedules: %v", err)
				}
			}
		}
	}

	log.Println("Schedules seeding completed")
	return nil
}

// getRandomClassCombination - Pilih kombinasi random dari available classes
// Bisa hanya Ekonomi, atau Ekonomi+Bisnis, atau Ekonomi+Bisnis+Eksekutif, dll
func getRandomClassCombination(availableClasses []string) model.StringArray {
	// Variasi kombinasi class yang mungkin
	combinations := [][]string{
		{"Ekonomi"},
		{"Bisnis"},
		{"Eksekutif"},
		{"Ekonomi", "Bisnis"},
		{"Ekonomi", "Eksekutif"},
		{"Bisnis", "Eksekutif"},
		{"Ekonomi", "Bisnis", "Eksekutif"},
	}

	// Pilih random kombinasi
	randIdx, err := faker.RandomInt(0, len(combinations)-1)
	if err != nil {
		// Fallback ke semua class jika error
		return model.StringArray(availableClasses)
	}

	return model.StringArray(combinations[randIdx[0]])
}

// SeedSchedules2026 - Tambah schedule khusus untuk Januari - Februari 2026
func SeedSchedules2026() error {
	var scheduleGroups []model.ScheduleGroup
	if err := pkg.DB.Find(&scheduleGroups).Error; err != nil {
		return err
	}

	var stations []model.Station
	if err := pkg.DB.Where("latitude IS NOT NULL AND longitude IS NOT NULL").Find(&stations).Error; err != nil {
		return err
	}

	if len(scheduleGroups) == 0 || len(stations) == 0 {
		log.Println("No schedule groups or stations found")
		return nil
	}

	log.Printf("Seeding additional schedules for Jan-Feb 2026 for %d schedule groups...", len(scheduleGroups))

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

		// Buat jadwal untuk Januari - Februari 2026 (setiap hari ada 1 keberangkatan)
		startDate := time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local)
		endDate := time.Date(2026, 2, 28, 23, 59, 59, 0, time.Local)

		for baseDate := startDate; baseDate.Before(endDate) || baseDate.Equal(endDate); baseDate = baseDate.AddDate(0, 0, 1) {
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
					ScheduleGroupID:    sg.ID,
					DepartureStationID: station.ID,
					ArrivalStationID:   station.ID,
					ArrivalTime:        arrivalTime,
					DepartureTime:      departureTime,
					Order:              order + 1,
					CreatedAt:          time.Now(),
					UpdatedAt:          time.Now(),
				}

				schedules = append(schedules, schedule)

				// Batch insert
				if len(schedules) >= batchSize {
					if err := pkg.DB.CreateInBatches(schedules, batchSize).Error; err != nil {
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
		if err := pkg.DB.CreateInBatches(schedules, batchSize).Error; err != nil {
			return err
		}
		log.Printf("Inserted final %d schedules", len(schedules))
	}

	log.Println("Jan-Feb 2026 schedules seeding completed")
	return nil
}
