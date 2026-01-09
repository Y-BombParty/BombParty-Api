package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/pkg/bomb"
	"bombparty.com/bombparty-api/pkg/game"
	"bombparty.com/bombparty-api/pkg/inventory"
	"bombparty.com/bombparty-api/pkg/team"
	"bombparty.com/bombparty-api/pkg/user"
)

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	// Initialisation des routes
	router := Routes(configuration)

	// Afficher toutes les routes
	printRoutes(router)

	log.Println("\nServing on : http://localhost:7774/api/v1/ \nServing swagger on : http://localhost:7774/swagger/index.html ")
	log.Fatal(http.ListenAndServe(":7774", router))
}

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:7774/swagger.json"),
	))

	// Serve Swagger JSON
	router.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.json")
	})

	router.Mount("/api/v1/bombs", bomb.Routes(configuration))
	router.Mount("/api/v1/user", user.Routes(configuration))
	router.Mount("/api/v1/inventory", inventory.Routes(configuration))
	router.Mount("/api/v1/games", game.Routes(configuration))
	router.Mount("/api/v1/teams", team.Routes(configuration))

	return router
}

func printRoutes(router chi.Router) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.ReplaceAll(route, "/*/", "/")
		log.Printf("%-6s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Printf("Logging err: %s\n", err.Error())
	}
}
