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

			// Handle %s banyak (field, param, value)
			switch verr.Tag() {
			case "before":
				messages = append(messages, fmt.Sprintf(format, verr.Field(), verr.Param()))
			case "min", "max":
				messages = append(messages, fmt.Sprintf(format, verr.Field(), verr.Param()))
			default:
				messages = append(messages, fmt.Sprintf(format, verr.Field()))
			}
		}
	}
	return messages
}

func (v *Validator) ValidateRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": []string{"Format data tidak valid"}})
		return false
	}
	msgs := v.ValidateStruct(req)
	if len(msgs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": msgs})
		return false
	}
	return true
}
