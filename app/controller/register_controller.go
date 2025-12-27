package controller

import (
	"TixTrain/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterController struct {
	Validator *pkg.Validator
}

func (ctl *RegisterController) Register(c *gin.Context) {
	var req RegisterRequest

	if !ctl.Validator.ValidateRequest(c, &req) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil"})
}
