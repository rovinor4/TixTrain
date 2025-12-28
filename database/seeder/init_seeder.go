package seeder

import "log"

func InitSeeder() {

	log.Println("Seeder Stations Started")
	if err := SeedStations(); err != nil {
		log.Fatal("Failed to seed stations:", err)
	}
	log.Println("Seeder Stations Completed")

	log.Println("Seeder Trains Started")
	if err := SeedTrains(); err != nil {
		log.Fatal("Failed to seed trains:", err)
	}

}
