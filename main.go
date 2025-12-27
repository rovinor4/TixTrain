package main

import (
	"TixTrain/app/controller"
	"TixTrain/database"
	"TixTrain/pkg"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

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
	err = database.ConnectDB()
	if err != nil {
		pkg.Logger.Error("Database connection failed", zap.Error(err))
		log.Fatal(err)
		return
	}

	if pkg.HandleCLI() {
		return
	}

	// Set Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Initialize Validator
	var validator pkg.Validator
	validator.InitValidator()

	regController := controller.RegisterController{Validator: &validator}

	r := gin.Default()
	SetupRoutes(r, &regController)

	// Gin Port
	ginPort := os.Getenv("GIN_PORT")
	pkg.Logger.Info("App started, listening on port 8080")
	err = r.Run(":" + ginPort)
	if err != nil {
		pkg.Logger.Error("Error starting server", zap.Error(err))
		return
	}

}
