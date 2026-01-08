package model

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type InventoryResponse struct {
	IDUser   uuid.UUID          `json:"id_user"`
	Elements []InventoryElement `json:"bombs"`
}

type InventoryElement struct {
	TypeBomb string `json:"type_bomb"`
	Amount   int    `json:"amount"`
}

type InventoryBombAmountChangePayload struct {
	Email    string `json:"email"`
	TypeBomb string `json:"type_bomb"`
	Amount   int    `json:"amount"`
}

func (i *InventoryBombAmountChangePayload) Bind(r *http.Request) error {
	if i.Email == "" {
		return errors.New("Email cannot be null")
	}
	if i.TypeBomb == "" {
		return errors.New("Type of the bomb cannot be null")
	}
	return nil
}
