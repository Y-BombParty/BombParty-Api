package model

import "net/http"

type BombRequest struct {
	Lat      float32 `json:"lat" binding:"required"`
	Long     float32 `json:"long" binding:"required"`
	TypeBomb string  `json:"type_bomb" binding:"required"`
	//IdUser   int     `json:"id_user" binding:"required"`
}

func (b *BombRequest) Bind(r *http.Request) error {
	return nil
}

type BombUpdateRequest struct {
	Lat      *float32 `json:"lat,omitempty"`
	Long     *float32 `json:"long,omitempty"`
	TypeBomb *string  `json:"type_bomb,omitempty"`
}

func (b *BombUpdateRequest) Bind(r *http.Request) error {
	return nil
}

type BombResponse struct {
	BombId   int     `json:"bomb_id"`
	Lat      float32 `json:"lat"`
	Long     float32 `json:"long"`
	TypeBomb string  `json:"type_bomb"`
	IdUser   int     `json:"id_user"`
}