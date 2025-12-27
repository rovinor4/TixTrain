package pkg

import (
	"TixTrain/database"
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
			log.Println("Seeding finished!")
			return true
		}
	}
	return false
}
