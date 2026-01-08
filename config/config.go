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
	GameRepository      dbmodel.GameRepository
	TeamRepository      dbmodel.TeamRepository
	BombRepository      dbmodel.BombRepository
}

func New() (*Config, error) {
	config := Config{
		JwtKey: os.Getenv("JWT_SECRET_KEY"),
		Port:   os.Getenv("PORT"),
	}

	databaseSession, err := gorm.Open(sqlite.Open("bomb-party.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}
	db.Migrate(databaseSession)

	config.UserRepository = dbmodel.NewUserRepository(databaseSession)
	config.InventoryRepository = dbmodel.NewInventoryRepository(databaseSession)
	config.GameRepository = dbmodel.NewGameRepository(databaseSession)
	config.TeamRepository = dbmodel.NewTeamRepository(databaseSession)
	config.BombRepository = dbmodel.NewBombRepository(databaseSession)
	return &config, nil
}
