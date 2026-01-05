package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"log"
	"math/rand"
	"time"
)

func SeedTickets() error {
	// Fetch necessary data
	var users []model.User
	var schedules []model.Schedule
	var coaches []model.Coach
	var stations []model.Station

	if err := pkg.DB.
		Joins("INNER JOIN user_roles ON user_roles.user_id = users.id").
		Joins("INNER JOIN roles ON roles.id = user_roles.role_id").
		Where("roles.name = ?", "passenger").
		Select("users.id").
		Find(&users).Error; err != nil {
		return err
	}
	if err := pkg.DB.Select("id, schedule_group_id").Preload("ScheduleGroup").Find(&schedules).Error; err != nil {
		return err
	}
	if err := pkg.DB.Select("id, class").Find(&coaches).Error; err != nil {
		return err
	}
	if err := pkg.DB.Select("id").Find(&stations).Error; err != nil {
		return err
	}

	if len(users) == 0 || len(schedules) == 0 || len(coaches) == 0 || len(stations) == 0 {
		log.Println("Missing required data for tickets")
		return nil
	}

	// Map coach classes for price determination
	coachClassMap := make(map[uint]string)
	for _, coach := range coaches {
		coachClassMap[coach.ID] = coach.Class
	}

	targetTickets := 500000
	batchSize := 2000

	log.Printf("Seeding %d tickets...", targetTickets)

	statuses := []string{"paid", "pending", "cancelled", "used"}
	statusWeights := []int{60, 20, 10, 10} // Weighted probabilities

	priceRanges := map[string][2]int64{
		"Ekonomi":   {50000, 150000},
		"Bisnis":    {150000, 300000},
		"Eksekutif": {300000, 600000},
	}

	for batch := 0; batch < targetTickets/batchSize; batch++ {
		tickets := make([]model.Ticket, 0, batchSize)

		for i := 0; i < batchSize; i++ {
			// Precompute random indices
			userIdx := rand.Intn(len(users))
			scheduleIdx := rand.Intn(len(schedules))
			schedule := schedules[scheduleIdx]

			// Random coach class for price determination
			coachIdx := rand.Intn(len(coaches))
			coachClass := coaches[coachIdx].Class
			priceRange := priceRanges[coachClass]
			price := rand.Int63n(priceRange[1]-priceRange[0]) + priceRange[0]

			// Weighted random status
			statusRand := rand.Intn(100) + 1
			var status string
			cumulative := 0
			for idx, weight := range statusWeights {
				cumulative += weight
				if statusRand <= cumulative {
					status = statuses[idx]
					break
				}
			}

			// Random created and updated dates
			daysAgo := rand.Intn(336) + 30 // 30-365 days ago
			createdAt := time.Now().AddDate(0, 0, -daysAgo)
			updatedAt := createdAt.AddDate(0, 0, rand.Intn(daysAgo))

			// Pilih departure dan arrival station random dari available stations
			departureStationIdx := rand.Intn(len(stations))
			arrivalStationIdx := rand.Intn(len(stations))

			// Jika sama, pilih lagi untuk arrival station
			for arrivalStationIdx == departureStationIdx {
				arrivalStationIdx = rand.Intn(len(stations))
			}

			ticket := model.Ticket{
				UserID:             users[userIdx].ID,
				ScheduleID:         schedule.ID,
				DepartureStationID: stations[departureStationIdx].ID,
				ArrivalStationID:   stations[arrivalStationIdx].ID,
				Price:              price,
				Status:             status,
				CreatedAt:          createdAt,
				UpdatedAt:          updatedAt,
			}

			tickets = append(tickets, ticket)
		}

		if err := pkg.DB.CreateInBatches(tickets, batchSize).Error; err != nil {
			return err
		}

		progress := (batch + 1) * batchSize
		percentage := float64(progress) / float64(targetTickets) * 100
		log.Printf("Progress: %d/%d tickets (%.2f%%)", progress, targetTickets, percentage)
	}

	log.Println("Tickets seeding completed - 500,000 tickets created!")
	return nil
}

// SeedTickets2026 - Seed tickets untuk schedule Jan-Feb 2026 agar ada quota yang terlihat
func SeedTickets2026() error {
	var users []model.User
	var schedules []model.Schedule
	var stations []model.Station

	// Ambil users
	if err := pkg.DB.
		Joins("INNER JOIN user_roles ON user_roles.user_id = users.id").
		Joins("INNER JOIN roles ON roles.id = user_roles.role_id").
		Where("roles.name = ?", "passenger").
		Select("users.id").
		Find(&users).Error; err != nil {
		return err
	}

	// Ambil schedules untuk Jan-Feb 2026
	startDate := time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2026, 2, 28, 23, 59, 59, 0, time.Local)

	if err := pkg.DB.Where("departure_time >= ? AND departure_time <= ?", startDate, endDate).
		Find(&schedules).Error; err != nil {
		return err
	}

	// Ambil stations
	if err := pkg.DB.Select("id").Find(&stations).Error; err != nil {
		return err
	}

	if len(users) == 0 || len(schedules) == 0 || len(stations) == 0 {
		log.Println("Missing required data for 2026 tickets")
		return nil
	}

	log.Printf("Seeding tickets for Jan-Feb 2026 schedules (%d schedules)...", len(schedules))

	statuses := []string{"paid", "used"}
	batchSize := 1000
	tickets := make([]model.Ticket, 0, batchSize)

	// Untuk setiap schedule, buat beberapa ticket agar ada quota yang terpakai
	for _, schedule := range schedules {
		// Random 1-5 tickets per schedule
		numTickets := rand.Intn(5) + 1

		for i := 0; i < numTickets; i++ {
			userIdx := rand.Intn(len(users))
			status := statuses[rand.Intn(len(statuses))]

			// Price berdasarkan class (50k-150k)
			price := rand.Int63n(100000) + 50000

			// Created time sebelum departure time
			daysBeforeDeparture := rand.Intn(30) + 1
			createdAt := schedule.DepartureTime.AddDate(0, 0, -daysBeforeDeparture)
			updatedAt := createdAt.Add(time.Duration(rand.Intn(24)) * time.Hour)

			// Pilih departure dan arrival station random
			departureStationIdx := rand.Intn(len(stations))
			arrivalStationIdx := rand.Intn(len(stations))

			// Jika sama, pilih lagi untuk arrival station
			for arrivalStationIdx == departureStationIdx {
				arrivalStationIdx = rand.Intn(len(stations))
			}

			ticket := model.Ticket{
				UserID:             users[userIdx].ID,
				ScheduleID:         schedule.ID,
				DepartureStationID: stations[departureStationIdx].ID,
				ArrivalStationID:   stations[arrivalStationIdx].ID,
				Price:              price,
				Status:             status,
				CreatedAt:          createdAt,
				UpdatedAt:          updatedAt,
			}

			tickets = append(tickets, ticket)

			// Batch insert
			if len(tickets) >= batchSize {
				if err := pkg.DB.CreateInBatches(tickets, batchSize).Error; err != nil {
					return err
				}
				log.Printf("Inserted %d tickets for 2026 schedules", len(tickets))
				tickets = tickets[:0]
			}
		}
	}

	// Insert remaining
	if len(tickets) > 0 {
		if err := pkg.DB.CreateInBatches(tickets, batchSize).Error; err != nil {
			return err
		}
		log.Printf("Inserted final %d tickets for 2026 schedules", len(tickets))
	}

	log.Println("Jan-Feb 2026 tickets seeding completed")
	return nil
}
