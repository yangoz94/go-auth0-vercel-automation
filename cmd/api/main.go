package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "auth0-vercel-script/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env.local file
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}

	// Create a new router
	router := chi.NewRouter()
	
	// Use the default middleware for logging each request
	router.Use(middleware.Logger)

	// Register the route(s)
	router.Post("/update-callback-urls", handlers.UpdateCallbackURLsHandler)

	// Here we go!
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}

