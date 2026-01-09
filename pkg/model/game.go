package model

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

type GameRequest struct {
	CenterLatitude  *float32   `json:"center_latitude"`
	CenterLongitude *float32   `json:"center_longitude"`
	Size            *float32   `json:"size"`
	StartingDate    *time.Time `json:"starting_date"`
	EndingDate      *time.Time `json:"ending_date"`
}

func (a *GameRequest) Bind(r *http.Request) error {

	minLimit := time.Now().AddDate(0, 0, 1) // 1 day
	maxLimit := time.Now().AddDate(0, 1, 0) // 1 month

	if a.CenterLatitude != nil {
		if *a.CenterLatitude < -90 || *a.CenterLatitude > 90 { // 90
			return errors.New("Wrong Center latitude value, must be betwen -90 / 90")
		}
	}

	if a.CenterLongitude != nil {
		if *a.CenterLongitude < -180 || *a.CenterLongitude > 180 {
			return errors.New("Wrong Center longitude value, must be between -180 / 180")
		}
	}

	if a.Size != nil {
		if *a.Size < 50 || *a.Size > 10107 {
			return errors.New("Wrong size value, must be between 50 and 10107")
		}
	}

	if a.EndingDate != nil {
		if a.EndingDate.After(maxLimit) || a.EndingDate.Before(minLimit) {
			return errors.New("Wrong ending date value, must be between 1 days and 1 month")
		}
	}

	return nil
}

func (a *GameRequest) ValidateCreate() error {

	missing := []string{}

	if a.CenterLatitude == nil {
		missing = append(missing, "center_latitude")
	}
	if a.CenterLongitude == nil {
		missing = append(missing, "center_longitude")
	}
	if a.Size == nil {
		missing = append(missing, "size")
	}
	if a.StartingDate == nil {
		missing = append(missing, "starting_date")
	}
	if a.EndingDate == nil {
		missing = append(missing, "ending_date")
	}

	// If there is missing fields, return all of theme
	if len(missing) > 0 {
		errorMsg := strings.Join(missing, ", ")
		return errors.New("Missing required fields: " + errorMsg)
	}

	return nil
}

func (a *GameRequest) ValidateUpdate() error {
	if a.CenterLatitude == nil &&
		a.CenterLongitude == nil &&
		a.Size == nil &&
		a.StartingDate == nil &&
		a.EndingDate == nil {
		return errors.New("At least one field must be provided")
	}
	return nil
}

type GameResponse struct {
	CenterLatitude  float32         `json:"center_latitude"`
	CenterLongitude float32         `json:"center_longitude"`
	Size            float32         `json:"size"`
	StartingDate    time.Time       `json:"starting_date"`
	EndingDate      time.Time       `json:"ending_date"`
	Teams           []*TeamResponse `json:"teams"`
}
