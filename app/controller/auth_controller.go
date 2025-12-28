package controller

import (
	"TixTrain/app/model"
	"TixTrain/database"
	"TixTrain/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	Validator *pkg.Validator
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ctl *AuthController) Login(c *gin.Context) {
	var req loginRequest

	if !ctl.Validator.ValidateRequest(c, &req) {
		return
	}

	var user model.User
	have := database.DB.Where("email = ?", req.Email).First(&user).RowsAffected

	if have == 0 || !pkg.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(401, gin.H{
			"message": "Email atau password salah",
			"error": gin.H{
				"email": "Email atau password salah",
			},
		})
		return
	}

	// make token
	GenerateToken, err := pkg.GenerateToken(20)
	if err != nil {
		c.JSON(500, gin.H{
			"message": pkg.GetMessage("error_server"),
		})

		pkg.Logger.Error("Error generating token", zap.Error(err))
		return
	}

	Token := model.Token{
		Value:     GenerateToken,
		UserID:    user.ID,
		UserAgent: c.Request.UserAgent(),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	err = database.DB.Create(&Token).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": pkg.GetMessage("error_server"),
		})
		pkg.Logger.Error("Error saving token", zap.Error(err))
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"data": gin.H{
			"token": GenerateToken,
			"user":  user,
		},
	})
}
