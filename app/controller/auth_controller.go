package controller

import (
	"TixTrain/app/model"
	"TixTrain/app/request"
	"TixTrain/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthController struct{}

//  NOTE: AuthController handles user authentication

func (a *AuthController) Login(c *gin.Context) {
	var req loginRequest

	if !pkg.GlobalValidator.ValidateRequest(c, &req) {
		return
	}

	var user model.User
	have := pkg.DB.Where("email = ?", req.Email).First(&user).RowsAffected

	if have == 0 || !new(pkg.Hash).CheckPasswordHash(req.Password, user.Password) {
		c.JSON(401, gin.H{
			"errors": gin.H{
				"email": "Email atau password salah",
			},
			"message": nil,
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

	err = pkg.DB.Create(&Token).Error
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
		},
	})
}

func (a *AuthController) Logout(c *gin.Context) {
	tokenString := c.GetString("token")

	err := pkg.DB.Where("value = ?", tokenString).Delete(&model.Token{}).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": pkg.GetMessage("error_server"),
		})
		pkg.Logger.Error("Error deleting token", zap.Error(err))
		return
	}

	c.JSON(200, gin.H{
		"message": "Logout successful",
	})

}

func (a *AuthController) Register(c *gin.Context) {

	var req request.RegisterRequest
	if !pkg.GlobalValidator.ValidateRequest(c, &req) {
		return
	}

	var errors map[string]string
	// Check if email already exists
	if pkg.DB.Model(&model.User{}).Where("email = ?", req.Email).RowsAffected > 0 {
		errors["email"] = "Email sudah terdaftar"
	}

	hashPassword, err := new(pkg.Hash).HashPassword(req.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": pkg.GetMessage("error_server"),
		})
		pkg.Logger.Error("Error hashing password", zap.Error(err))
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
		Role:     "passenger",
	}

	err = pkg.DB.Create(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": pkg.GetMessage("error_server"),
		})
		pkg.Logger.Error("Error creating user", zap.Error(err))
		return
	}

	// make identity card
	dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid date format for DateOfBirth",
		})
		return
	}

	identityCard := model.IdentityCard{
		Type:        req.IdentityType,
		Number:      req.IdentityCardNumber,
		Name:        req.Name,
		Gender:      req.Gender,
		DateOfBirth: &dateOfBirth,
		IsMe:        true,
		UserID:      user.ID,
	}

	// Save identityCard to the database
	err = pkg.DB.Create(&identityCard).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": pkg.GetMessage("error_server"),
		})
		pkg.Logger.Error("Error creating identity card", zap.Error(err))
		return
	}

	c.JSON(201, gin.H{
		"message": "Registration successful",
	})
}
