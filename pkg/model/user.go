package model

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type UserCreatePayload struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserCreatePayload) Bind(r *http.Request) error {
	if u.UserName == "" {
		return errors.New("UserName cannot be null")
	}
	if u.Email == "" {
		return errors.New("Email cannot be null")
	}
	if u.Password == "" {
		return errors.New("Password cannot be null")
	}
	return nil
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserLoginPayload) Bind(r *http.Request) error {
	if u.Email == "" {
		return errors.New("Email cannot be null")
	}
	if u.Password == "" {
		return errors.New("Password cannot be null")
	}
	return nil
}

type UserUpdatePayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

func (u *UserUpdatePayload) Bind(r *http.Request) error {
	return nil
}

type UserResponse struct {
	IdUser   uuid.UUID `json:"id_user"`
	Email    string    `json:"email"`
	UserName string    `json:"user_name"`
	IdTeam   uuid.UUID `json:"id_team"`
}
