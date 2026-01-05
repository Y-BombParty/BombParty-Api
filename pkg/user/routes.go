package user

import (
	"bombparty.com/bombparty-api/config"
	"github.com/go-chi/chi/v5"
)

func Routes(config *config.Config) *chi.Mux {
	userConfig := New(config)

	router := chi.NewRouter()

	router.Post("/login", userConfig.Login)
	router.Post("/register", userConfig.Register)
	return router
}
