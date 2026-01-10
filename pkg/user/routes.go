package user

import (
	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/pkg/authentication"
	"github.com/go-chi/chi/v5"
)

func ProtectedRoutes(config *config.Config) *chi.Mux {
	UserConfig := New(config)

	router := chi.NewRouter()

	router.Use(authentication.AuthMiddleware(config.JwtKey))
	router.Put("/update", UserConfig.Update)
	router.Delete("/delete", UserConfig.DeleteUser)
	router.Get("/user", UserConfig.GetOneUser)
	return router
}
