package main

import (
	"TixTrain/database"
	"TixTrain/database/seeder"
	"TixTrain/pkg"
	"TixTrain/route"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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

func main() {

	// Initialize Logger
	pkg.InitLog()
	defer func(Logger *zap.Logger) {
		_ = Logger.Sync()
	}(pkg.Logger)

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		pkg.Logger.Error("Error loading .env file")
		log.Fatal("Error loading .env file")
	}

	// Connect to Database
	err = pkg.ConnectDB()
	if err != nil {
		pkg.Logger.Error("Database connection failed", zap.Error(err))
		log.Fatal(err)
		return
	}

	if HandleCLI() {
		return
	}

	// Set Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Initialize Global Validator
	pkg.InitValidator()

	// Gin Setup
	r := gin.Default()
	r.Use(pkg.SetupCORS())
	route.SetupRoutes(r)

	// Start server
	ginPort := os.Getenv("GIN_PORT")
	pkg.Logger.Info("App started, listening on port 8080")
	err = r.Run(":" + ginPort)
	if err != nil {
		pkg.Logger.Error("Error starting server", zap.Error(err))
		return
	}
}
