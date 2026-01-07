package game

import (
	"bombparty.com/bombparty-api/config"
	//"bombparty.com/bombparty-api/pkg/authentication"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {

	// Init Router
	gameConfig := New(configuration)
	router := chi.NewRouter()

	// Routes protected by authentication
	router.Group(func(router chi.Router) {
		//router.Use(authentication.AuthMiddleware(""))

		router.Get("/", gameConfig.GetAlldHandler)
		router.Get("/{id}", gameConfig.GetByIdHandler)
		router.Post("/", gameConfig.PostHandler)
		router.Put("/{id}", gameConfig.UpdateHandler)
		router.Delete("/{id}", gameConfig.DeleteHandler)
	})

	return router
}
