package models

import "github.com/google/uuid"

type InventoryResponse struct {
	IDUser   uuid.UUID          `json:"id_user"`
	Elements []InventoryElement `json:"bombs"`
}

type InventoryElement struct {
	TypeBomb string `json:"type_bomb"`
	Amount   int    `json:"amoun"`
}
