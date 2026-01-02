package route

import (
	"TixTrain/app/controller"
	"TixTrain/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	AuthMiddleware := middleware.AuthMiddleware()
	RoleAdmin := middleware.RoleMiddleware("admin")
	//RolePassenger := middleware.RoleMiddleware("passenger")
	//RoleStaff := middleware.RoleMiddleware("staff")

	{
		auth := r.Group("/auth")
		auth.POST("/register", new(controller.RegisterController).Register)
		auth.POST("/login", new(controller.AuthController).Login)
		auth.GET("/logout", new(controller.AuthController).Logout).Use(AuthMiddleware)
	}

	{
		station := r.Group("/stations")
		station.GET("/list", new(controller.StationController).Get)
		station.GET("/show/:id", new(controller.StationController).Show)
		station.POST("/create", new(controller.StationController).Create).Use(AuthMiddleware, RoleAdmin)
		station.POST("/update/:id", new(controller.StationController).Update).Use(AuthMiddleware, RoleAdmin)
		station.DELETE("/delete/:id", new(controller.StationController).Delete).Use(AuthMiddleware, RoleAdmin)
	}
}
