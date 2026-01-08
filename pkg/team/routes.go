package team

import (
	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/pkg/authentication"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	teamConfig := New(configuration)
	router := chi.NewRouter()

	router.Use(authentication.AuthMiddleware(configuration.JwtKey))
	router.Post("/", teamConfig.CreateTeamHandler)
	router.Get("/", teamConfig.GetAllTeamsHandler)
	router.Get("/{id}", teamConfig.GetTeamByIDHandler)
	router.Put("/{id}", teamConfig.UpdateTeamHandler)
	router.Delete("/{id}", teamConfig.DeleteTeamHandler)

	return router
}
