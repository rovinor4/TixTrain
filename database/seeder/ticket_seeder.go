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

	if err := pkg.DB.Select("id").Where("role", "passenger").Find(&users).Error; err != nil {
		return err
	}
	if err := pkg.DB.Select("id").Find(&schedules).Error; err != nil {
		return err
	}
	if err := pkg.DB.Select("id, class").Find(&coaches).Error; err != nil {
		return err
	}

	if len(users) == 0 || len(schedules) == 0 || len(coaches) == 0 {
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

			ticket := model.Ticket{
				UserID:     users[userIdx].ID,
				ScheduleID: schedules[scheduleIdx].ID,
				Price:      price,
				Status:     status,
				CreatedAt:  createdAt,
				UpdatedAt:  updatedAt,
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
