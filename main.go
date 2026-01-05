package main

import (
	"fmt"
	"log"
	"net/http"

	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/pkg/inventory"
	"bombparty.com/bombparty-api/pkg/user"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func Routes(config *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/api/v1/user", user.Routes(config))
	router.Mount("/api/v1/inventory", inventory.Routes(config))
	return router
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	config, _ := config.New()

	router := Routes(config)

	log.Println("Serving on ", config.Port)
	log.Fatal(http.ListenAndServe(":7774", router))
}
