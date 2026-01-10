package authentication

import (
	"bombparty.com/bombparty-api/config"
	"github.com/go-chi/chi/v5"
)

func Routes(config *config.Config) *chi.Mux {
	authConfig := New(config)

	router := chi.NewRouter()

	router.Post("/login",authConfig.Login)
	router.Post("/register", authConfig.Register)

	return router
}