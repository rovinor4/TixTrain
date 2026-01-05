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

	// 3. Seed Roles (diperlukan untuk role based access control)
	log.Println("=== Seeder Roles Started ===")
	if err := SeedRoles(); err != nil {
		log.Fatal("Failed to seed roles:", err)
	}
	log.Println("=== Seeder Roles Completed ===\n")

	// 3b. Seed Permissions (diperlukan untuk role permissions)
	log.Println("=== Seeder Permissions Started ===")
	if err := SeedPermissions(); err != nil {
		log.Fatal("Failed to seed permissions:", err)
	}
	log.Println("=== Seeder Permissions Completed ===\n")

	// 3c. Seed Role Permissions (menghubungkan roles dengan permissions)
	log.Println("=== Seeder Role Permissions Started ===")
	if err := SeedRolePermissions(); err != nil {
		log.Fatal("Failed to seed role permissions:", err)
	}
	log.Println("=== Seeder Role Permissions Completed ===\n")

	// 4. Seed Users (diperlukan untuk tickets dan tokens)
	log.Println("=== Seeder Users Started ===")
	if err := SeedUsers(); err != nil {
		log.Fatal("Failed to seed users:", err)
	}
	log.Println("=== Seeder Users Completed ===\n")

	// 4b. Seed Identity Cards (diperlukan untuk ticket details)
	log.Println("=== Seeder Identity Cards Started ===")
	if err := SeedIdentityCards(); err != nil {
		log.Fatal("Failed to seed identity cards:", err)
	}
	log.Println("=== Seeder Identity Cards Completed ===\n")

	// 5. Seed Schedules & Schedule Groups (diperlukan untuk coaches dan tickets)
	log.Println("=== Seeder Schedules Started ===")
	if err := SeedSchedules(); err != nil {
		log.Fatal("Failed to seed schedules:", err)
	}
	log.Println("=== Seeder Schedules Completed ===\n")

	// 6. Seed Coaches (diperlukan untuk seats, bergantung pada schedule groups)
	log.Println("=== Seeder Coaches Started ===")
	if err := SeedCoaches(); err != nil {
		log.Fatal("Failed to seed coaches:", err)
	}
	log.Println("=== Seeder Coaches Completed ===\n")

	// 7. Seed Seats (diperlukan untuk tickets)
	log.Println("=== Seeder Seats Started ===")
	if err := SeedSeats(); err != nil {
		log.Fatal("Failed to seed seats:", err)
	}
	log.Println("=== Seeder Seats Completed ===\n")

	// 8. Seed Tickets (TARGET: 500,000 records)
	log.Println("=== Seeder Tickets Started (Target: 500,000) ===")
	if err := SeedTickets(); err != nil {
		log.Fatal("Failed to seed tickets:", err)
	}
	log.Println("=== Seeder Tickets Completed ===\n")

	// 8b. Seed Tickets for Jan-Feb 2026 (tambahan untuk testing dengan quota)
	log.Println("=== Seeder Tickets Jan-Feb 2026 Started ===")
	if err := SeedTickets2026(); err != nil {
		log.Fatal("Failed to seed tickets 2026:", err)
	}
	log.Println("=== Seeder Tickets Jan-Feb 2026 Completed ===\n")

	// 9. Seed Ticket Details (diperlukan untuk relasi ticket-seat-identity)
	log.Println("=== Seeder Ticket Details Started ===")
	if err := SeedTicketDetails(); err != nil {
		log.Fatal("Failed to seed ticket details:", err)
	}
	log.Println("=== Seeder Ticket Details Completed ===\n")

	log.Println("🎉 ALL SEEDING COMPLETED SUCCESSFULLY! 🎉")
	log.Println("Total tickets: 500,000")
}
