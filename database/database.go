package database

import (
	"log"

	"bombparty.com/bombparty-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init the DB, create and migrate
func InitDatabase() {

	var err error
	DB, err = gorm.Open(sqlite.Open("bomb-party.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	Migrate(DB)

	log.Println("Database connected")
}

// Init the DB by migrate evrey models
func Migrate(db *gorm.DB) {

	db.AutoMigrate(
		&dbmodel.GameEntry{},
	)

	log.Println("Database migrated successfully")
}
