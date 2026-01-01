package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"fmt"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
)

func SeedSeats() error {
	var coaches []model.Coach
	if err := pkg.DB.Find(&coaches).Error; err != nil { // Change from pkg.DB to database.DB
		return err
	}

	if len(coaches) == 0 {
		log.Println("No coaches found")
		return nil
	}

	log.Printf("Seeding seats for %d coaches...", len(coaches))

	batchSize := 1000
	seats := make([]model.Seat, 0, batchSize)

	for _, coach := range coaches {
		// Jumlah kursi berdasarkan kelas
		var numSeats int
		switch coach.Class {
		case "Ekonomi":
			numSeats = 64 // 4 kursi per baris, 16 baris
		case "Bisnis":
			numSeats = 48 // 4 kursi per baris, 12 baris
		case "Eksekutif":
			numSeats = 40 // 4 kursi per baris, 10 baris
		default:
			numSeats = 50
		}

		// Generate nomor kursi (1A, 1B, 1C, 1D, 2A, 2B, dst)
		seatLetters := []string{"A", "B", "C", "D"}
		seatNumber := 1

		for i := 0; i < numSeats; i++ {
			letter := seatLetters[i%4]
			if i > 0 && i%4 == 0 {
				seatNumber++
			}

			daysAgoSlice, err := faker.RandomInt(30, 365)
			if err != nil {
				return err
			}
			daysAgo := daysAgoSlice[0]

			updatedDaysAgoSlice, err := faker.RandomInt(0, 30)
			if err != nil {
				return err
			}
			updatedDaysAgo := updatedDaysAgoSlice[0]

			seat := model.Seat{
				CoachID:   coach.ID,
				Number:    fmt.Sprintf("%d%s", seatNumber, letter),
				CreatedAt: time.Now().AddDate(0, 0, -daysAgo),
				UpdatedAt: time.Now().AddDate(0, 0, -updatedDaysAgo),
			}

			seats = append(seats, seat)

			// Batch insert
			if len(seats) >= batchSize {
				if err := pkg.DB.CreateInBatches(seats, batchSize).Error; err != nil {
					return err
				}
				log.Printf("Inserted %d seats", len(seats))
				seats = seats[:0]
			}
		}
	}

	// Insert remaining
	if len(seats) > 0 {
		if err := pkg.DB.CreateInBatches(seats, batchSize).Error; err != nil {
			return err
		}
		log.Printf("Inserted final %d seats", len(seats))
	}

	log.Println("Seats seeding completed")
	return nil
}
