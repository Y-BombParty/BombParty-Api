package dbmodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BombEntry struct {
	BombID   uuid.UUID `gorm:"type:uuid; primaryKey"`
	Lat      float32   `json:"lat"`
	Long     float32   `json:"long"`
	TypeBomb string    `json:"type_bomb"`
	IdUser   int       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"care_taker"`
}

type BombRepository interface {
	Create(bomb *BombEntry) (*BombEntry, error)
	FindAll() ([]*BombEntry, error)
	FindAllByUserId(userId int) ([]*BombEntry, error)
	FindById(id int) (*BombEntry, error)
	Update(bomb *BombEntry) (*BombEntry, error)
	Delete(id int) error
}

type bombRepository struct {
	db *gorm.DB
}

func NewBombRepository(db *gorm.DB) BombRepository {
	return &bombRepository{db: db}
}

func (r *bombRepository) Create(bomb *BombEntry) (*BombEntry, error) {
	if err := r.db.Create(bomb).Error; err != nil {
		return nil, err
	}
	return bomb, nil
}

func (r *bombRepository) FindAll() ([]*BombEntry, error) {
	var bombs []*BombEntry
	if err := r.db.Find(&bombs).Error; err != nil {
		return nil, err
	}
	return bombs, nil
}

func (r *bombRepository) FindAllByUserId(userId int) ([]*BombEntry, error) {
	var bombs []*BombEntry
	if err := r.db.Where("type_bomb = ?", userId).Find(&bombs).Error; err != nil {
		return nil, err
	}

	return bombs, nil
}

func (r *bombRepository) FindById(id int) (*BombEntry, error) {
	var bomb BombEntry
	if err := r.db.Find(&bomb, id).Error; err != nil {
		return nil, err
	}
	return &bomb, nil
}

func (r *bombRepository) Update(bomb *BombEntry) (*BombEntry, error) {
	if err := r.db.Save(bomb).Error; err != nil {
		return nil, err
	}
	return bomb, nil
}

func (r *bombRepository) Delete(id int) error {
	if err := r.db.Delete(&BombEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}
