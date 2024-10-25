package main

import (
	"fmt"
	"log"
	"net/http"
	"spoonacular-api/db"
	"spoonacular-api/routes"
)

func main() {
	// Initialize the database connection
	err := db.InitDB("localhost", "5432", "myuser", "password", "spoonacularapidb")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.DB.Close()

	// Set up routes
	routes.SetupUserRoutes()
	routes.SetupAuthRoutes()
	routes.SetupRecipeRoutes()

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
