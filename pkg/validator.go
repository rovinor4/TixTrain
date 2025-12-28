package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate      *validator.Validate
	MessageErrors map[string]string
}

func (v *Validator) InitValidator() {
	v.validate = validator.New()
	v.MessageErrors = map[string]string{
		"required": "%s wajib di isi",
		"before":   "%s harus sebelum %s",
		"min":      "%s minimal %s karakter",
		"max":      "%s maksimal %s karakter",
		"email":    "%s harus berupa email yang valid",
	}
}

func (v *Validator) ValidateStruct(data interface{}) []string {
	var messages []string
	err := v.validate.Struct(data)
	if err != nil {
		for _, verr := range err.(validator.ValidationErrors) {
			format, ok := v.MessageErrors[verr.Tag()]
			if !ok {
				format = "%s tidak valid"
			}

			if verr.Param() != "" {
				messages = append(messages, fmt.Sprintf(format, verr.Field(), verr.Param()))
			} else {
				messages = append(messages, fmt.Sprintf(format, verr.Field()))
			}
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
			"error":   inputError,
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
