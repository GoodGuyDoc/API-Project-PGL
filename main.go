package main

import (
	"fmt"
	"log"
	"net/http"
	"spoonacular-api/db"
	"spoonacular-api/routes"
	"testing"
)

func setupServer() error {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	_, err := http.Get("http://localhost:8080/")
	return err
}

func main() {
	// Initialize the database connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.DB.Close()

	// Set up routes
	go routes.SetupUserRoutes()
	go routes.SetupAuthRoutes()
	go routes.SetupRecipeRoutes()

	setupServer()
	// Set up static file server
}

func TestSetupServer(t *testing.T) {
	var isFailed = setupServer()
	if isFailed != nil {
		t.Fatalf("Test Setup failed. No other aspects of the program can continue...")
	}
	t.Log("TestSetupServer Completed")
}
