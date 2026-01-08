package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameEntry struct {
	IDGame          uuid.UUID `gorm:"type:uuid;primaryKey"`
	CenterLatitude  float32   `json:"center_latitude"`
	CenterLongitude float32   `json:"center_longitude"`
	Size            float32   `json:"size"`
	StartingDate    time.Time `json:"starting_date"`
	EndingDate      time.Time `json:"ending_date"`

	Teams []TeamEntry `json:"teams" gorm:"foreignKey:IDTeam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CrudInfo
}

func (g *GameEntry) BeforeCreate(tx *gorm.DB) (err error) {
	g.IDGame = uuid.New()
	return
}

type GameRepository interface {
	Create(entry *GameEntry) (*GameEntry, error)
	FindById(id uuid.UUID) (*GameEntry, error)
	FindAll() ([]*GameEntry, error)
	Update(entry *GameEntry, id uuid.UUID) (*GameEntry, error)
	DeleteById(id uuid.UUID) error
}

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return &gameRepository{db: db}
}

func (r *gameRepository) Create(entry *GameEntry) (*GameEntry, error) {

	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}

	return entry, nil
}

func (r *gameRepository) FindById(id uuid.UUID) (*GameEntry, error) {

	var entries *GameEntry
	if err := r.db.Model(&GameEntry{}).
		Preload("Teams").
		First(&entries, id).Error; err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *gameRepository) FindAll() ([]*GameEntry, error) {

	var entries []*GameEntry
	if err := r.db.Model(&GameEntry{}).
		Preload("Teams").
		Find(&entries).Error; err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *gameRepository) Update(entry *GameEntry, id uuid.UUID) (*GameEntry, error) {

	result := r.db.Model(&GameEntry{}).
		Preload("Teams").
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"centerlatitude":  entry.CenterLatitude,
			"centerlongitude": entry.CenterLongitude,
			"size":            entry.Size,
			"startingdate":    entry.StartingDate,
			"endingdate":      entry.EndingDate,
		})

	if result.Error != nil {
		return nil, result.Error
	}

	// Check if something has been update
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return entry, nil

}

func (r *gameRepository) DeleteById(id uuid.UUID) error {

	if err := r.db.Delete(GameEntry{}, id).Error; err != nil {
		return err
	}

	return nil
}
