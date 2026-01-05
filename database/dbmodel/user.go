package dbmodel

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntry struct {
	IDUser   uuid.UUID `gorm:"type:uuid; primaryKey"`
	UserName string    `gorm:"type:varchar(255);unique"`
	Email    string    `gorm:"type:varchar(255);unique"`
	Password string
	IDTeam   *uuid.UUID `gorm:"type:uuid"`
	Team     *TeamEntry `gorm:"foreignKey:IDTeam;references:IDTeam"`

	CrudInfo
}

type UserRepository interface {
	Register(entry *UserEntry) (*UserEntry, error)
	FindOne(filter, value string) (*UserEntry, error)
	FindAll() ([]*UserEntry, error)
	Login(entry *UserEntry) (*UserEntry, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]*UserEntry, error) {
	var entries []*UserEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *userRepository) FindOne(filter, value string) (*UserEntry, error) {
	var entries []*UserEntry
	if err := r.db.Where(filter+" = ?", value).Find(&entries).Error; err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, errors.New("No entry found with the filter : " + filter + " and the value : " + value)
	}
	if len(entries) > 1 {
		return nil, errors.New("Too much entries found with the filter : " + filter + " and the value : " + value)
	}
	return entries[0], nil
}

func (r *userRepository) Register(entry *UserEntry) (*UserEntry, error) {
	entry.IDUser = uuid.New()
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *userRepository) Login(entry *UserEntry) (*UserEntry, error) {
	user, err := r.FindOne("email", entry.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("No user found with this email")
	}

	hashedPassword := user.Password
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(entry.Password)) != nil {
		return nil, errors.New("Invalid email or password")
	}

	return user, nil
}
