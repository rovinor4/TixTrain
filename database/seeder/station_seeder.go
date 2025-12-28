package seeder

import (
	"TixTrain/app/model"
	"TixTrain/database"
	"encoding/json"
	"os"
)

type StationJSON struct {
	Name      string   `json:"name"`
	Code      string   `json:"code"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}

func SeedStations() error {
	data, err := os.ReadFile("database/seeder/json/stations.json")
	if err != nil {
		return err
	}

	// Parse JSON
	var stationsJSON []StationJSON
	if err := json.Unmarshal(data, &stationsJSON); err != nil {
		return err
	}

	// Convert ke model Station dan insert
	for _, sJSON := range stationsJSON {
		station := model.Station{
			Name: sJSON.Name,
			Code: sJSON.Code,
		}

		// Handle nullable latitude/longitude
		if sJSON.Latitude != nil {
			station.Latitude = *sJSON.Latitude
		}
		if sJSON.Longitude != nil {
			station.Longitude = *sJSON.Longitude
		}

		if err := database.DB.Where("code = ?", station.Code).FirstOrCreate(&station).Error; err != nil {
			return err
		}
	}

	return nil
}
