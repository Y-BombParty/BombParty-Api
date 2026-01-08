package model

import (
	"errors"
	"net/http"
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
