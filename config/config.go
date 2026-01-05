package config

import (
	"bombparty.com/bombparty-api/database/dbmodel"
	"bombparty.com/bombparty-api/database"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	BombRepository       dbmodel.BombRepository
}

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)

	config.BombRepository = dbmodel.NewBombRepository(databaseSession)

	return &config, nil
}
