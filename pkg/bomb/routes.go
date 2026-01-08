package bomb

import (
	"bombparty.com/bombparty-api/config"
	//"bombparty.com/bombparty-api/pkg/authentication"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	bombConfig := New(configuration)
	router := chi.NewRouter()
	
	router.Group(func(r chi.Router) {
		//r.Use(authentication.AuthMiddleware("your_secret_key"))
		
		// Create
		r.Post("/", bombConfig.CreateBomb)
		
		// Read
		r.Get("/", bombConfig.GetAllBombs)
		r.Get("/{id}", bombConfig.GetBomb)
		r.Get("/user/{userId}", bombConfig.GetBombsByUserId)
		
		// Update
		r.Put("/{id}", bombConfig.UpdateBomb)
		
		// Delete
		r.Delete("/{id}", bombConfig.DeleteBomb)
	})
	
	return router
}