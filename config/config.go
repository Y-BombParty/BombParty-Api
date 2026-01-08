package config

import (
	"bombparty.com/bombparty-api/database"
	"bombparty.com/bombparty-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	// Repository connection
	GameRepository dbmodel.GameRepository
	TeamRepository dbmodel.TeamRepository
	BombRepository dbmodel.BombRepository
}

func New() (*Config, error) {

	config := Config{}

	// Init DB connection
	databaseSession, err := gorm.Open(sqlite.Open("bomb-party.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	// Models migrate
	database.Migrate(databaseSession)

	// Init repository
	config.GameRepository = dbmodel.NewGameRepository(databaseSession)
	config.TeamRepository = dbmodel.NewTeamRepository(databaseSession)
	config.BombRepository = dbmodel.NewBombRepository(databaseSession)

	return &config, nil
}
