package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	db "bombparty.com/bombparty-api/database"
	"bombparty.com/bombparty-api/database/dbmodel"
)

type Config struct {
	Port                string
	JwtKey              string
	UserRepository      dbmodel.UserRepository
	InventoryRepository dbmodel.InventoryRepository
}

func New() (*Config, error) {
	config := Config{
		JwtKey: os.Getenv("JWT_SECRET_KEY"),
		Port:   os.Getenv("PORT"),
	}

	databaseSession, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	db.Migrate(databaseSession)

	config.UserRepository = dbmodel.NewUserRepository(databaseSession)
	config.InventoryRepository = dbmodel.NewInventoryRepository(databaseSession)
	return &config, nil
}
