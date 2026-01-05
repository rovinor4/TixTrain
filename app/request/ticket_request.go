package request

type TicketRequestFind struct {
	DepartureStation string `json:"departure_station" form:"departure_station" validate:"required"`
	ArrivalStation   string `json:"arrival_station" form:"arrival_station" validate:"required"`
	DepartureDate    string `json:"departure_date" form:"departure_date" validate:"required,datetime=2006-01-02"`
	CountPassenger   int    `json:"count_passenger" form:"count_passenger" validate:"required,min=1"`
	Classes          string `json:"classes" form:"classes"`
}
