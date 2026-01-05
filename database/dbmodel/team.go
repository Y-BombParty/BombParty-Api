package dbmodel

import (
	"github.com/google/uuid"
)

type TeamEntry struct {
	IDTeam uuid.UUID `gorm:"type:uuid;primaryKey"`
	Score  int       `gorm:"type:integer;"`
	Name   string    `gorm:"type:varchar(255);"`
	Color  string    `gorm:"type:varchar(255);"`

	CrudInfo
}
