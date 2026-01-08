package inventory

import (
	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/pkg/authentication"
	"github.com/go-chi/chi/v5"
)

func Routes(config *config.Config) *chi.Mux {
	inventoryConfig := New(config)

	router := chi.NewRouter()
	router.Use(authentication.AuthMiddleware(config.JwtKey))
	router.Get("/init", inventoryConfig.InitUserInventory)
	router.Get("/inventory", inventoryConfig.GetUserInventory)
	router.Post("/add", inventoryConfig.ChangeBombsAmount)
	return router
}
