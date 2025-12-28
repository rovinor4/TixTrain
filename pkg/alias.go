package pkg

import (
	"TixTrain/database"
	"TixTrain/database/seeder"
	"log"
	"os"
	"strings"
)

func HandleCLI() bool {
	if len(os.Args) > 1 {
		cmd := strings.ToLower(os.Args[1])
		switch cmd {
		case "migrate":
			database.Migrate()
			return true
		case "seed":
			log.Println("Seeding...")
			seeder.InitSeeder()
			log.Println("Seeding finished!")
			return true
		}
	}
	return false
}
