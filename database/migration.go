package database

import (
	"TixTrain/app/model"
	"log"
)

func Migrate() {
	log.Println("Migrating...")

	//err := DB.Migrator().DropTable(
	//	&model.User{},
	//	&model.Token{},
	//	&model.Station{},
	//	&model.Train{},
	//	&model.ScheduleGroup{},
	//	&model.Schedule{},
	//	&model.Coach{},
	//	&model.Seat{},
	//	&model.Ticket{},
	//)
	//if err != nil {
	//	return
	//}

	err := DB.AutoMigrate(
		&model.User{},
		&model.Token{},
		&model.Station{},
		&model.Train{},
		&model.ScheduleGroup{},
		&model.Schedule{},
		&model.Coach{},
		&model.Seat{},
		&model.Ticket{},
	)
	if err != nil {
		return
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Migration finished!")
}
