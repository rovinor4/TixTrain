package controller

import (
	"TixTrain/app/model"
	"TixTrain/app/request"
	"TixTrain/pkg"
	_ "TixTrain/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StationController struct {
}

func (s *StationController) Get(c *gin.Context) {
	var Stations []model.Station
	var total int64

	search := c.DefaultQuery("q", "")
	withOutId := c.DefaultQuery("withoutId", "")

	if withOutId != "" {
		if _, err := strconv.Atoi(withOutId); err != nil {
			c.JSON(400, gin.H{"errors": "Parameter 'withoutId' harus berupa angka"})
			return
		}
	}

	scopeFunc, page, pageSize, offset := pkg.Paginate(c, 10)
	OnDatabase := pkg.DB.Scopes(scopeFunc)
	if search != "" {
		OnDatabase.Where("name ILIKE ?", "%"+search+"%")
	}
	if withOutId != "" {
		OnDatabase.Where("id != ?", withOutId)
	}
	OnDatabase.Find(&Stations).Count(&total)

	for i := range Stations {
		Stations[i].Name = new(pkg.Helper).TitleCase(Stations[i].Name)
	}

	c.JSON(200, gin.H{
		"data":      Stations,
		"page":      page,
		"page_size": pageSize,
		"offset":    offset,
		"total":     total,
	})
}

func (s *StationController) Create(c *gin.Context) {
	var req request.StationRequest
	if !pkg.GlobalValidator.ValidateRequest(c, &req) {
		return
	}

	errInput := make(map[string]string)
	var existingStation model.Station
	if err := pkg.DB.Where("code = ?", req.Code).First(&existingStation).Error; err == nil {
		errInput["code"] = "Kode stasiun sudah digunakan"
	}

	if len(errInput) > 0 {
		c.JSON(400, gin.H{"errors": errInput})
		return
	}

	station := model.Station{
		Name:      req.Name,
		Code:      req.Code,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	if err := pkg.DB.Create(&station).Error; err != nil {
		c.JSON(500, gin.H{"message": "Gagal membuat stasiun"})
		pkg.Logger.Error("Failed Create Station", zap.Error(err))
		return
	}

	c.JSON(201, gin.H{"data": station})
}

func (s *StationController) Show(c *gin.Context) {
	id := c.Param("id")
	var station model.Station
	if err := pkg.DB.First(&station, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "Stasiun tidak ditemukan"})
		return
	}

	c.JSON(200, gin.H{"data": station})
}

func (s *StationController) Update(c *gin.Context) {
	id := c.Param("id")
	var station model.Station
	if err := pkg.DB.First(&station, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "Stasiun tidak ditemukan"})
		return
	}

	var req request.StationRequest
	if !pkg.GlobalValidator.ValidateRequest(c, &req) {
		return
	}

	errInput := make(map[string]string)
	var existingStation model.Station
	if err := pkg.DB.Where("code = ? AND id != ?", req.Code, id).First(&existingStation).Error; err == nil {
		errInput["code"] = "Kode stasiun sudah digunakan"
	}

	if len(errInput) > 0 {
		c.JSON(400, gin.H{"errors": errInput})
		return
	}

	station.Name = req.Name
	station.Code = req.Code
	station.Longitude = req.Longitude
	station.Latitude = req.Latitude

	if err := pkg.DB.Save(&station).Error; err != nil {
		c.JSON(500, gin.H{"message": "Gagal memperbarui stasiun"})
		pkg.Logger.Error("Failed Update Station", zap.Error(err))
		return
	}

	c.JSON(200, gin.H{"data": station})
}

func (s *StationController) Delete(c *gin.Context) {
	id := c.Param("id")
	var station model.Station
	if err := pkg.DB.First(&station, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "Stasiun tidak ditemukan"})
		return
	}

	if err := pkg.DB.Delete(&station).Error; err != nil {
		c.JSON(500, gin.H{"message": "Gagal menghapus stasiun"})
		pkg.Logger.Error("Failed Delete Station", zap.Error(err))
		return
	}

	c.JSON(200, gin.H{"message": "Stasiun berhasil dihapus"})
}
