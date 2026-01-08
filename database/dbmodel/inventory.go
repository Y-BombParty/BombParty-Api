package dbmodel

import (
	"slices"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryEntry struct {
	IDUser   uuid.UUID `gorm:"type:uuid;"`
	User     UserEntry `gorm:"foreignKey:IDUser; references:IDUser"`
	TypeBomb string
	Amount   int

	CrudInfo
}

type InventoryRepository interface {
	FindAll() ([]*InventoryEntry, error)
	FindByUser(user UserEntry) ([]*InventoryEntry, error)
	ChangeBombsAmount(user UserEntry, typeBomb string, amount int) (*InventoryEntry, error)
	AddNewBombType(user UserEntry, typeBomb string, startingAmount int) (*InventoryEntry, error)
	InitUserInventory(user UserEntry) ([]*InventoryEntry, error)
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) FindAll() ([]*InventoryEntry, error) {
	var entries []*InventoryEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *inventoryRepository) FindByUser(user UserEntry) ([]*InventoryEntry, error) {
	var entries []*InventoryEntry
	if err := r.db.Where("id_user = ?", user.IDUser).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *inventoryRepository) ChangeBombsAmount(user UserEntry, typeBomb string, amount int) (*InventoryEntry, error) {
	userBombs, err := r.FindByUser(user)
	if err != nil {
		return nil, err
	}

	idx := slices.IndexFunc(userBombs, func(b *InventoryEntry) bool {
		return b.TypeBomb == typeBomb
	})
	userBombs[idx].Amount = max(userBombs[idx].Amount+amount, 0)
	if err = r.db.Where("id_user = ? AND type_bomb = ?", user.IDUser, typeBomb).UpdateColumns(userBombs[idx]).Error; err != nil {
		return nil, err
	}
	return userBombs[idx], nil
}

func (r *inventoryRepository) AddNewBombType(user UserEntry, typeBomb string, startingAmount int) (*InventoryEntry, error) {
	userBombs, err := r.FindByUser(user)
	if err != nil {
		return nil, err
	}

	for _, ele := range userBombs {
		if typeBomb == ele.TypeBomb {
			ele.Amount = startingAmount
			if err = r.db.Where("id_user = ? AND type_bomb = ?", user.IDUser, typeBomb).UpdateColumns(ele).Error; err != nil {
				return nil, err
			}
			return nil, nil
		}
	}

	bombType := &InventoryEntry{
		IDUser:   user.IDUser,
		TypeBomb: typeBomb,
		Amount:   0,
	}
	if err := r.db.Create(bombType).Error; err != nil {
		return nil, err
	}
	return bombType, nil
}

func (r *inventoryRepository) InitUserInventory(user UserEntry) ([]*InventoryEntry, error) {
	if _, err := r.AddNewBombType(user, "classic", 0); err != nil {
		return nil, err
	}
	if _, err := r.AddNewBombType(user, "double", 0); err != nil {
		return nil, err
	}
	if _, err := r.AddNewBombType(user, "giant", 0); err != nil {
		return nil, err
	}
	return r.FindByUser(user)
}
