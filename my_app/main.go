// main.go
package main

import (
	"net/http"
	"log"
	"my_app/router"
)

func main() {
	// Create a new router
	appRouter := router.NewRouter()

	// Configure routes
	appRouter.SetupRoutes()

	// Set up a static file server for serving CSS, JS, and images
	appRouter.GetHandler().PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the HTTP server
	log.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", appRouter.GetHandler()))
}

