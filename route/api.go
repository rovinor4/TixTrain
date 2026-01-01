package route

import (
	"TixTrain/app/controller"
	"TixTrain/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	AuthMiddleware := middleware.AuthMiddleware()

	{
		auth := r.Group("/auth")
		auth.POST("/register", new(controller.RegisterController).Register)
		auth.POST("/login", new(controller.AuthController).Login)
		auth.GET("/logout", new(controller.AuthController).Logout).Use(AuthMiddleware)
	}

	{
		station := r.Group("/stations")
		station.GET("/list", new(controller.StationController).Get)
	}
}
