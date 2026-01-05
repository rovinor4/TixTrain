package controller

import (
	"TixTrain/app/model"
	"TixTrain/app/request"
	"TixTrain/pkg"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ScheduleController struct {
}

func (s *ScheduleController) FindSchedule(c *gin.Context) {

	// 1. parse request
	var req request.TicketRequestFind
	if err := c.ShouldBindQuery(&req); err != nil {
		pkg.Logger.Error("Failed to bind query", zap.Error(err))
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	if req.ArrivalStation == req.DepartureStation {
		c.JSON(400, gin.H{"errors": "Stasiun keberangkatan dan stasiun kedatangan tidak boleh sama"})
		return
	}

	// formating DepartureDate to YYYY-MM-DD for querying
	departureDate, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		pkg.Logger.Error("Failed to parse departure date", zap.Error(err))
		c.JSON(400, gin.H{"errors": "Invalid departure date format"})
		return
	}

	var StationDeparture model.Station
	pkg.DB.Where("id = ?", req.DepartureStation).First(&StationDeparture)
	if StationDeparture.ID == 0 {
		pkg.Logger.Error("Departure station not found", zap.String("station", req.DepartureStation))
		c.JSON(400, gin.H{"errors": "Departure station not found"})
		return
	}

	var StationArrival model.Station
	pkg.DB.Where("id = ?", req.ArrivalStation).First(&StationArrival)
	if StationArrival.ID == 0 {
		pkg.Logger.Error("Arrival station not found", zap.String("station", req.ArrivalStation))
		c.JSON(400, gin.H{"errors": "Arrival station not found"})
		return
	}

	var scheduleGroups []model.ScheduleGroup
	var respond []map[string]any
	query := pkg.DB.
		Where("departure_stations @> ?", `[`+fmt.Sprint(req.DepartureStation)+`]`).
		Where("arrival_stations @> ?", `[`+fmt.Sprint(req.ArrivalStation)+`]`).
		Where("departure_time >= ? AND departure_time <= ?",
			departureDate.Format("2006-01-02 15:04:05-07:00"),
			departureDate.Add(24*time.Hour).Format("2006-01-02 15:04:05-07:00"))

	if req.Classes != "" {
		query = query.Where("class @> ?", `["`+req.Classes+`"]`)
	}

	if err := query.Order("departure_time ASC").
		Preload("Schedules", func(db *gorm.DB) *gorm.DB {
			return db.Where("departure_station_id = ? OR arrival_station_id = ?", StationDeparture.ID, StationArrival.ID).
				Order("\"order\" asc")
		}).
		Preload("Train").
		Find(&scheduleGroups).Error; err != nil {
		pkg.Logger.Error("Failed to find schedules", zap.Error(err))
		c.JSON(500, gin.H{"errors": "Failed to fetch schedules"})

	}

	type Quota struct {
		SumData int64
	}
	for _, scheduleGroup := range scheduleGroups {
		var quota Quota
		pkg.DB.Model(&model.Coach{}).Select("SUM(quota) as sum_data").Where("schedule_group_id = ?", scheduleGroup.ID).Scan(&quota)

		if quota.SumData >= int64(req.CountPassenger) && len(scheduleGroup.Schedules) >= 2 {
			tickitResponse := map[string]interface{}{
				"id":             scheduleGroup.ID,
				"name":           scheduleGroup.Name,
				"class":          scheduleGroup.Class,
				"departure_time": scheduleGroup.Schedules[0].DepartureTime.Format("02-01-2006 15:04"),
				"arrival_time":   scheduleGroup.Schedules[1].ArrivalTime.Format("02-01-2006 15:04"),
				"train":          scheduleGroup.Train,
				"quota":          quota.SumData,
			}
			respond = append(respond, tickitResponse)
		}

	}

	if StationDeparture.Image != nil {
		image := new(pkg.Helper).Assets(*StationDeparture.Image)
		StationDeparture.Image = &image
	}

	if StationArrival.Image != nil {
		image := new(pkg.Helper).Assets(*StationArrival.Image)
		StationArrival.Image = &image
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"station_departure": StationDeparture,
			"station_arrival":   StationArrival,
			"schedule_groups":   respond,
		},
	})
}
