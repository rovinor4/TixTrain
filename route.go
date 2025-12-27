package main

import (
	"TixTrain/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, regController *controller.RegisterController) {
	r.POST("/register", regController.Register)
}
