package database

import (
	"log"

	"bombparty.com/bombparty-api/database/dbmodel"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.UserEntry{},
		&dbmodel.TeamEntry{},
		&dbmodel.InventoryEntry{},
	)
	log.Println("Database migratted succesfuly")
}
