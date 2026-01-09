package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"bombparty.com/bombparty-api/config"
	_ "bombparty.com/bombparty-api/docs" // Import pour initialiser Swagger
	"bombparty.com/bombparty-api/pkg/bomb"
	"bombparty.com/bombparty-api/pkg/game"
	"bombparty.com/bombparty-api/pkg/inventory"
	"bombparty.com/bombparty-api/pkg/team"
	"bombparty.com/bombparty-api/pkg/user"
	"github.com/joho/godotenv"
)

// @title BombParty API
// @version 1.0
// @description API pour le jeu BombParty
// @host localhost:7774
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" suivi de votre token JWT

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Valeur par défaut
	}

	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	// Initialisation des routes
	router := Routes(configuration, PORT)

	// Afficher toutes les routes
	printRoutes(router)

	log.Printf("\nServing on: http://localhost%s/api/v1/", PORT)
	log.Printf("Serving swagger on: http://localhost%s/swagger/index.html\n", PORT)
	
	address := fmt.Sprintf("%s", PORT)
	log.Fatal(http.ListenAndServe(address, router))
}

func Routes(configuration *config.Config, port string) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// URL Swagger dynamique basée sur le port
	swaggerURL := fmt.Sprintf("http://localhost%s/swagger/doc.json", port)
	
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
	))

	// Serve Swagger JSON
	router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
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