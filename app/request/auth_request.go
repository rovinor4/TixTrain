package request

type RegisterRequest struct {
	Name               string `json:"name" validate:"required"`
	Email              string `json:"email" validate:"required,email"`
	Password           string `json:"password" validate:"required,min=8"`
	PasswordConfirm    string `json:"password_confirm" validate:"required"`
	DateOfBirth        string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	Gender             string `json:"gender" validate:"required,oneof=Male Female"`
	IdentityType       string `json:"identity_type" validate:"required,oneof=KTP SIM Passport"`
	IdentityCardNumber string `json:"identity_card_number" validate:"required"`
}
