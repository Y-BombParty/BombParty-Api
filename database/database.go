package database

import (
	"bombparty.com/bombparty-api/database/dbmodel"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.Bomb{},
	)
	log.Println("Database migrated successfully")
}
