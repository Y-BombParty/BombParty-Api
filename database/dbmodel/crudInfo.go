package dbmodel

import "time"

type CrudInfo struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `sql:"index" json:"deleted_at"`
}