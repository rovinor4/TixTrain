package pkg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate      *validator.Validate
	MessageErrors map[string]string
}

var GlobalValidator *Validator

func InitValidator() {
	GlobalValidator = &Validator{
		validate: validator.New(),
		MessageErrors: map[string]string{
			"required": "%s wajib di isi",
			"before":   "%s harus sebelum %s",
			"min":      "%s minimal %s karakter",
			"max":      "%s maksimal %s karakter",
			"email":    "%s harus berupa email yang valid",
			"datetime": "%s harus berupa tanggal yang valid dengan format %s",
			"oneof":    "%s harus salah satu dari nilai berikut: %s",
		},
	}
}

func (v *Validator) ValidateStruct(data interface{}) map[string]string {
	messages := make(map[string]string)
	err := v.validate.Struct(data)
	if err != nil {
		for _, verr := range err.(validator.ValidationErrors) {
			format, ok := v.MessageErrors[verr.Tag()]
			if !ok {
				format = "%s tidak valid"
			}

			var message string
			if verr.Param() != "" {
				message = fmt.Sprintf(format, verr.Field(), verr.Param())
			} else {
				message = fmt.Sprintf(format, verr.Field())
			}

			// Convert field name ke lowercase untuk key
			fieldName := strings.ToLower(verr.Field())
			messages[fieldName] = message
		}
	}
	return messages
}

func (v *Validator) ValidateRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Ada kesalahan, format json tidak valid.",
		})
		return false
	}
	inputError := v.ValidateStruct(req)
	if len(inputError) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Ada kesalahan, check kembali.",
			"errors":  inputError,
		})
		return false
	}
	return true
}

func GetMessage(tag string) string {
	template := map[string]string{
		"error_server": "Terjadi kesalahan pada server",
	}

	if v, ok := template[tag]; ok {
		return v
	}

	return "Tidak ada pesan error"
}
