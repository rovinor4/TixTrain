package request

type StationRequest struct {
	Name      string  `json:"name" validate:"required"`
	Code      string  `json:"code" validate:"required,min=3,max=10"`
	Longitude float64 `json:"longitude" validate:"required,min=0"`
	Latitude  float64 `json:"latitude" validate:"required,min=0"`
}
