package controller

import (
	"TixTrain/app/model"
	"TixTrain/pkg"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TicketController struct {
}

func (t *TicketController) GetTicketDetail(c *gin.Context) {
	id := c.Param("id")
	class := c.Param("class")

	var scheduleGroup model.ScheduleGroup
	if err := pkg.DB.Where("id = ?", id).
		Find(&scheduleGroup).Error; err != nil {

		pkg.Logger.Error("Failed to find schedule group by ID", zap.Error(err))
		c.JSON(500, gin.H{"errors": "Failed to fetch schedule group"})
		return
	}

	if scheduleGroup.ID == 0 {
		c.JSON(404, gin.H{"errors": "Schedule group not found"})
		return
	}

	var coachesRespond []map[string]any
	var coaches []model.Coach
	if err := pkg.DB.Where("schedule_group_id = ?", scheduleGroup.ID).Where("class = ?", class).Preload("Seats").
		Find(&coaches).Error; err != nil {
		c.JSON(500, gin.H{"errors": "Failed to fetch coaches"})
		return
	}

	// Kumpulkan semua seat IDs dari semua coaches
	seatIDs := make([]uint, 0)
	for _, coach := range coaches {
		for _, seat := range coach.Seats {
			seatIDs = append(seatIDs, seat.ID)
		}
	}

	// Query sekali untuk semua seats yang booked
	bookedSeats := make(map[uint]bool)
	if len(seatIDs) > 0 {
		var tickets []model.TicketDetail
		if err := pkg.DB.Where("seat_id IN ?", seatIDs).Select("seat_id").Find(&tickets).Error; err != nil {
			c.JSON(500, gin.H{"errors": "Failed to fetch tickets"})
			return
		}

		// Buat map untuk quick lookup
		for _, ticket := range tickets {
			bookedSeats[ticket.ID] = true
		}
	}

	// Build response dengan data dari map (no more queries)
	for _, coach := range coaches {
		seatsRespond := []map[string]any{}
		for _, seat := range coach.Seats {
			seatR := map[string]interface{}{
				"id":        seat.ID,
				"number":    seat.Number,
				"is_booked": bookedSeats[seat.ID],
			}
			seatsRespond = append(seatsRespond, seatR)
		}
		coachR := map[string]interface{}{
			"id":    coach.ID,
			"class": coach.Class,
			"price": coach.Price,
			"code":  coach.Code,
			"seats": seatsRespond,
		}

		coachesRespond = append(coachesRespond, coachR)
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"schedule_group": scheduleGroup,
			"coaches":        coachesRespond,
		},
	})
}
