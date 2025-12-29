package seeder

import (
	"TixTrain/app/model"
	"TixTrain/database"
	"log"
	"math/rand"
)

func SeedTicketDetails() error {
	// Fetch all tickets
	var tickets []model.Ticket
	if err := database.DB.Select("id, user_id, created_at, updated_at").Find(&tickets).Error; err != nil {
		return err
	}

	if len(tickets) == 0 {
		log.Println("No tickets found for ticket details")
		return nil
	}

	// Fetch all seats
	var seats []model.Seat
	if err := database.DB.Select("id").Find(&seats).Error; err != nil {
		return err
	}

	if len(seats) == 0 {
		log.Println("No seats found for ticket details")
		return nil
	}

	// Fetch all identity cards grouped by user
	var identityCards []model.IdentityCard
	if err := database.DB.Select("id, user_id").Find(&identityCards).Error; err != nil {
		return err
	}

	if len(identityCards) == 0 {
		log.Println("No identity cards found for ticket details")
		return nil
	}

	// Map identity cards by user ID for quick lookup
	identityCardsByUser := make(map[uint][]uint)
	for _, ic := range identityCards {
		identityCardsByUser[ic.UserID] = append(identityCardsByUser[ic.UserID], ic.ID)
	}

	batchSize := 2000
	targetDetails := len(tickets) // Each ticket gets 1 ticket detail

	log.Printf("Seeding %d ticket details...", targetDetails)

	ticketDetails := make([]model.TicketDetail, 0, batchSize)

	for idx, ticket := range tickets {
		// Get identity cards for this user
		userIdentityCards, hasIdentityCards := identityCardsByUser[ticket.UserID]

		// Skip if user has no identity cards
		if !hasIdentityCards || len(userIdentityCards) == 0 {
			continue
		}

		// Random seat
		seatIdx := rand.Intn(len(seats))

		// Random identity card from user's identity cards
		identityCardIdx := rand.Intn(len(userIdentityCards))

		ticketDetail := model.TicketDetail{
			TicketID:       ticket.ID,
			SeatID:         seats[seatIdx].ID,
			IdentityCardID: userIdentityCards[identityCardIdx],
			CreatedAt:      ticket.CreatedAt,
			UpdatedAt:      ticket.UpdatedAt,
		}

		ticketDetails = append(ticketDetails, ticketDetail)

		// Batch insert
		if len(ticketDetails) >= batchSize || idx == len(tickets)-1 {
			if len(ticketDetails) > 0 {
				if err := database.DB.CreateInBatches(ticketDetails, batchSize).Error; err != nil {
					return err
				}
				progress := idx + 1
				percentage := float64(progress) / float64(len(tickets)) * 100
				log.Printf("Progress: %d/%d ticket details (%.2f%%)", progress, len(tickets), percentage)
				ticketDetails = make([]model.TicketDetail, 0, batchSize)
			}
		}
	}

	log.Println("Ticket details seeding completed")
	return nil
}
