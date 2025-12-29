package seeder

import "log"

func InitSeeder() {
	// 1. Seed Stations (diperlukan untuk schedules)
	log.Println("=== Seeder Stations Started ===")
	if err := SeedStations(); err != nil {
		log.Fatal("Failed to seed stations:", err)
	}
	log.Println("=== Seeder Stations Completed ===\n")

	// 2. Seed Trains (diperlukan untuk coaches dan schedule groups)
	log.Println("=== Seeder Trains Started ===")
	if err := SeedTrains(); err != nil {
		log.Fatal("Failed to seed trains:", err)
	}
	log.Println("=== Seeder Trains Completed ===\n")

	// 3. Seed Users (diperlukan untuk tickets dan tokens)
	log.Println("=== Seeder Users Started ===")
	if err := SeedUsers(); err != nil {
		log.Fatal("Failed to seed users:", err)
	}
	log.Println("=== Seeder Users Completed ===\n")

	// 3b. Seed Identity Cards (diperlukan untuk ticket details)
	log.Println("=== Seeder Identity Cards Started ===")
	if err := SeedIdentityCards(); err != nil {
		log.Fatal("Failed to seed identity cards:", err)
	}
	log.Println("=== Seeder Identity Cards Completed ===\n")

	// 4. Seed Schedule Groups & Coaches (diperlukan untuk schedules dan seats)
	log.Println("=== Seeder Schedule Groups & Coaches Started ===")
	if err := SeedScheduleGroupsAndCoaches(); err != nil {
		log.Fatal("Failed to seed schedule groups and coaches:", err)
	}
	log.Println("=== Seeder Schedule Groups & Coaches Completed ===\n")

	// 5. Seed Schedules (diperlukan untuk tickets)
	log.Println("=== Seeder Schedules Started ===")
	if err := SeedSchedules(); err != nil {
		log.Fatal("Failed to seed schedules:", err)
	}
	log.Println("=== Seeder Schedules Completed ===\n")

	// 6. Seed Seats (diperlukan untuk tickets)
	log.Println("=== Seeder Seats Started ===")
	if err := SeedSeats(); err != nil {
		log.Fatal("Failed to seed seats:", err)
	}
	log.Println("=== Seeder Seats Completed ===\n")

	// 7. Seed Tickets (TARGET: 500,000 records)
	log.Println("=== Seeder Tickets Started (Target: 500,000) ===")
	if err := SeedTickets(); err != nil {
		log.Fatal("Failed to seed tickets:", err)
	}
	log.Println("=== Seeder Tickets Completed ===\n")

	// 8. Seed Ticket Details (diperlukan untuk relasi ticket-seat-identity)
	log.Println("=== Seeder Ticket Details Started ===")
	if err := SeedTicketDetails(); err != nil {
		log.Fatal("Failed to seed ticket details:", err)
	}
	log.Println("=== Seeder Ticket Details Completed ===\n")

	log.Println("🎉 ALL SEEDING COMPLETED SUCCESSFULLY! 🎉")
	log.Println("Total tickets: 500,000")
}
