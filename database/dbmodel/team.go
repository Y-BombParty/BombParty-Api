package dbmodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamEntry struct {
	IDTeam uuid.UUID `gorm:"type:uuid;primaryKey"`
	Score  int       `gorm:"type:integer;"`
	Name   string    `gorm:"type:varchar(255);"`
	Color  string    `gorm:"type:varchar(255);"`
	IDGame uuid.UUID
	CrudInfo
}

func (b *TeamEntry) BeforeCreate(tx *gorm.DB) (err error) {
	b.IDTeam = uuid.New()
	return
}

type TeamRepository interface {
	Create(team *TeamEntry) (*TeamEntry, error)
	FindAll() ([]*TeamEntry, error)
	FindById(uuid uuid.UUID) (*TeamEntry, error)
	Update(team *TeamEntry) (*TeamEntry, error)
	Delete(uuid uuid.UUID, team *TeamEntry) error
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) Create(team *TeamEntry) (*TeamEntry, error) {
	if err := r.db.Create(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

func (r *teamRepository) Delete(uuid uuid.UUID, team *TeamEntry) error {
	return r.db.Delete(team, uuid).Error
}

func (r *teamRepository) FindById(uuid uuid.UUID) (*TeamEntry, error) {
	var team TeamEntry
	if err := r.db.First(&team, uuid).Error; err != nil {
		return nil, err
	}
	return &team, nil

}

func (r *teamRepository) Update(team *TeamEntry) (*TeamEntry, error) {
	if err := r.db.Save(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

func (r *teamRepository) FindAll() ([]*TeamEntry, error) {
	var teams []*TeamEntry
	if err := r.db.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
