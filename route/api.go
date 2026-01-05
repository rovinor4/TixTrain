package route

import (
	"TixTrain/app/controller"
	"TixTrain/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	AuthMiddleware := middleware.AuthMiddleware()

	r.Static("/assets", "./storage/public")
	{
		auth := r.Group("/auth")
		auth.POST("/register", new(controller.RegisterController).Register)
		auth.POST("/login", new(controller.AuthController).Login)
		auth.GET("/logout", new(controller.AuthController).Logout).Use(AuthMiddleware)
	}

	{
		admin := r.Group("/admin")
		admin.Use(middleware.AuthPermissionMiddleware())
		{
			admin.GET("/stats", new(controller.DashboardController).GetStats)
			admin.GET("/tickets", new(controller.DashboardController).GetTickets)
			admin.GET("/tickets/all", new(controller.DashboardController).GetTicketsAll)
			admin.GET("/stations", new(controller.StationController).Get)
		}
	}

	// Staff routes
	{
		staff := r.Group("/staff")
		staff.Use(middleware.AuthPermissionMiddleware())
		{
			staff.GET("/tickets", new(controller.DashboardController).GetTickets)
		}
	}

	// Passenger routes
	{
		passenger := r.Group("/passenger")
		passenger.Use(middleware.AuthPermissionMiddleware())
		{
			passenger.GET("/tickets", new(controller.DashboardController).GetTickets)
		}
	}

	{
		station := r.Group("/stations")
		station.GET("/list", new(controller.StationController).Get)
		station.GET("/show/:id", new(controller.StationController).Show)
		station.POST("/create", new(controller.StationController).Create).Use(AuthMiddleware)
		station.POST("/update/:id", new(controller.StationController).Update).Use(AuthMiddleware)
		station.DELETE("/delete/:id", new(controller.StationController).Delete).Use(AuthMiddleware)
	}

	{
		ticket := r.Group("/schedules")
		ticket.GET("/find", new(controller.ScheduleController).FindSchedule)
	}

	{
		ticket := r.Group("/tickets")
		{
			bill := ticket.Group("/bill")
			bill.GET("/detail/:id/:class", new(controller.TicketController).GetTicketDetail)
		}
	}
}
