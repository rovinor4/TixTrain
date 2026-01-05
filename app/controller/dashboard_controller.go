package controller

import (
	"TixTrain/app/model"
	"TixTrain/pkg"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

// GetTickets - Dengan pagination (Scenario 1)
func (d *DashboardController) GetTickets(c *gin.Context) {
	var tickets []model.Ticket
	var total int64

	scopeFunc, page, pageSize, offset := pkg.Paginate(c, 10)

	pkg.DB.Model(&model.Ticket{}).
		Select("id, user_id, schedule_id, price, status, created_at").
		Scopes(scopeFunc).
		Find(&tickets)

	pkg.DB.Model(&model.Ticket{}).Count(&total)

	c.JSON(200, gin.H{
		"data":      tickets,
		"page":      page,
		"page_size": pageSize,
		"offset":    offset,
		"total":     total,
	})
}

// GetTicketsAll - Tanpa pagination (Scenario 2)
func (d *DashboardController) GetTicketsAll(c *gin.Context) {
	var tickets []model.Ticket

	pkg.DB.Select("id, user_id, schedule_id, price, status, created_at").
		Find(&tickets)

	c.JSON(200, gin.H{
		"data":  tickets,
		"total": len(tickets),
	})
}

// GetStats - Summary
func (d *DashboardController) GetStats(c *gin.Context) {
	var stats struct {
		TotalUsers    int64 `json:"total_users"`
		TotalTickets  int64 `json:"total_tickets"`
		TotalStations int64 `json:"total_stations"`
		TotalTrains   int64 `json:"total_trains"`
		TotalRevenue  int64 `json:"total_revenue"`
	}

	pkg.DB.Model(&model.User{}).Count(&stats.TotalUsers)
	pkg.DB.Model(&model.Ticket{}).Count(&stats.TotalTickets)
	pkg.DB.Model(&model.Station{}).Count(&stats.TotalStations)
	pkg.DB.Model(&model.Train{}).Count(&stats.TotalTrains)
	pkg.DB.Model(&model.Ticket{}).Select("COALESCE(SUM(price), 0)").Scan(&stats.TotalRevenue)

	c.JSON(200, gin.H{"data": stats})
}
