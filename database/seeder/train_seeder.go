package seeder

import (
	"TixTrain/app/model"
	"TixTrain/database"
	"encoding/json"
	"os"
)

type TrainJSON struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func SeedTrains() error {
	data, err := os.ReadFile("database/seeder/json/trains.json")
	if err != nil {
		return err
	}

	// Parse JSON
	var trainsJSON []TrainJSON
	if err := json.Unmarshal(data, &trainsJSON); err != nil {
		return err
	}

	// Convert ke model Train dan insert
	for _, tJSON := range trainsJSON {
		train := model.Train{
			Name: tJSON.Name,
			Code: tJSON.Code,
		}

		if err := database.DB.Where("code = ?", train.Code).FirstOrCreate(&train).Error; err != nil {
			return err
		}
	}

	return nil
}
