package dbmodel

import (
	"gorm.io/gorm"
)

type Bomb struct {
	BombID   uint    `gorm:"primaryKey;autoIncrement;column:bomb_id" json:"bomb_id"`
	Lat      float32 `json:"lat"`
	Long     float32 `json:"long"`
	TypeBomb string  `json:"type_bomb"`
	// IdUser   int     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"care_taker"`
}

type BombRepository interface {
	Create(bomb *Bomb) (*Bomb, error)
	FindAll() ([]*Bomb, error)
	FindAllByUserId(userId int) ([]*Bomb, error)
	Find(id int) (*Bomb, error)
	Update(bomb *Bomb) (*Bomb, error)
	Delete(id int) error
}


type bombRepository struct {
	db *gorm.DB
}

func NewBombRepository(db *gorm.DB) BombRepository {
	return &bombRepository{db: db}
}

func (r *bombRepository) Create(bomb *Bomb) (*Bomb, error) {
	if err := r.db.Create(bomb).Error; err != nil {
		return nil, err
	}
	return bomb, nil
}

func (r *bombRepository) FindAll() ([]*Bomb, error) {
	var bombs []*Bomb
	if err := r.db.Find(&bombs).Error; err != nil {
		return nil, err
	}
	return bombs, nil
}

func (r *bombRepository) FindAllByUserId(userId int) ([]*Bomb, error) {
	var bombs []*Bomb
	if err := r.db.Where("type_bomb = ?", userId).Find(&bombs).Error; err != nil {
		return nil, err
	}

	return bombs, nil
}

func (r *bombRepository) Find(id int) (*Bomb, error) {
	var bomb Bomb
	if err := r.db.Find(&bomb, id).Error; err != nil {
		return nil, err
	}
	return &bomb, nil
}

func (r *bombRepository) Update(bomb *Bomb) (*Bomb, error) {
	if err := r.db.Save(bomb).Error; err != nil {
		return nil, err
	}
	return bomb, nil
}

func (r *bombRepository) Delete(id int) error {
	if err := r.db.Delete(&Bomb{}, id).Error; err != nil {
		return err
	}
	return nil
}
