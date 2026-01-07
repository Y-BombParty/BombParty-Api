package model

import (
	"errors"
	"net/http"
	"time"
)

type GameRequest struct {
	CenterLatitude  float32   `json:"game_center_latitude"`
	CenterLongitude float32   `json:"game_center_longitude"`
	Size            float32   `json:"game_size"`
	StartingDate    time.Time `json:"game_starting_date"`
	EndingDate      time.Time `json:"game_ending_date`
}

func (a *GameRequest) Bind(r *http.Request) error {

	if a.CenterLatitude < -90 || a.CenterLatitude > 90 { // 90
		return errors.New("Wrong Center latitude value, must be betwen -90 / 90")
	}

	if a.CenterLongitude < -180 || a.CenterLongitude > 180 {
		return errors.New("Wrong Center longitude value, must be between -180 / 180")
	}

	if a.Size < 50 || a.Size > 10107 {
		return errors.New("Wrong size value, must be between 50 and 10107")
	}

	return nil
}

type GameResponse struct {
	CenterLatitude  float32         `json:"game_center_latitude"`
	CenterLongitude float32         `json:"game_center_longitude"`
	Size            float32         `json:"game_size"`
	StartingDate    time.Time       `json:"game_starting_date"`
	EndingDate      time.Time       `json:"game_ending_date`
	Teams           []*TeamResponse `json:"game_teams"`
}
