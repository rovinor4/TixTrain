package main

import (
	"TixTrain/app/controller"
	"TixTrain/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", new(controller.RegisterController).Register)
	r.POST("/login", new(controller.AuthController).Login)

	auth := r.Group("")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/logout", new(controller.AuthController).Logout)

	}
}
