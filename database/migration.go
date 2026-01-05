package database

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"log"
)

func Migrate() {
	log.Println("Migrating...")

	err := pkg.DB.Migrator().DropTable(
		&model.RolePermission{},
		&model.Role{},
		&model.Permission{},
		&model.UserRole{},
		&model.User{},
		&model.IdentityCard{},
		&model.Token{},
		&model.Station{},
		&model.Train{},
		&model.ScheduleGroup{},
		&model.Schedule{},
		&model.Coach{},
		&model.Seat{},
		&model.Ticket{},
		&model.TicketDetail{},
	)
	if err != nil {
		return
	}

	err = pkg.DB.AutoMigrate(
		&model.Role{},
		&model.Permission{},
		&model.RolePermission{},
		&model.User{},
		&model.UserRole{},
		&model.IdentityCard{},
		&model.Token{},
		&model.Station{},
		&model.Train{},
		&model.ScheduleGroup{},
		&model.Schedule{},
		&model.Coach{},
		&model.Seat{},
		&model.Ticket{},
		&model.TicketDetail{},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Migration finished!")
}
