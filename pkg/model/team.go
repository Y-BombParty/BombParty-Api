package model

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type TeamRequest struct {
	Score  int       `json:"score"`
	Name   string    `json:"name"`
	Color  string    `json:"color"`
	IDGame uuid.UUID `json:"id_game"`
}

func (t *TeamRequest) Bind(r *http.Request) error {
	if t.Score < 0 {
		return errors.New("The score must not be negative")
	}
	if t.Name == "" {

		return errors.New("The name must not be null")
	}
	if t.Color == "" {

		return errors.New("The color must not be null")
	}
	if t.IDGame.String() == "" {

		return errors.New("The id_game must not be null")
	}

	return nil

}

type TeamResponse struct {
	Score  int       `json:"score"`
	Name   string    `json:"name"`
	Color  string    `json:"color"`
	IDGame uuid.UUID `json:"id_game"`
}
