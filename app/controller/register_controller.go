package controller

import (
	"TixTrain/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterController struct{}

func (ctl *RegisterController) Register(c *gin.Context) {
	var req registerRequest

	if !pkg.GlobalValidator.ValidateRequest(c, &req) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil"})
}
