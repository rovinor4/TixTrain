package controller

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	_ "TixTrain/pkg"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
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
			c.JSON(400, gin.H{"error": "Parameter 'withoutId' harus berupa angka"})
			return
		}
	}

	scopeFunc, page, pageSize, offset := pkg.Paginate(c, 10)
	pkg.DB.Scopes(scopeFunc).Where("name ILIKE ?", "%"+search+"%").Where("id != ?", withOutId).Find(&Stations).Count(&total)

	for i := range Stations {
		Stations[i].Name = fmt.Sprintf("Stasiun %s", new(pkg.Helper).TitleCase(Stations[i].Name))
	}

	c.JSON(200, gin.H{
		"data":      Stations,
		"page":      page,
		"page_size": pageSize,
		"offset":    offset,
		"total":     total,
	})
}
