package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	handlers "auth0-vercel-script/internal/handlers"
)

func main() {
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

