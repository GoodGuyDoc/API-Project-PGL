package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"spoonacular-api/api"
	"spoonacular-api/db"
	"spoonacular-api/session"
)

func SetupUserRoutes() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/profile", ProfilePageHandler)
	http.HandleFunc("/api/profile", ProfileHandler)
	http.HandleFunc("/api/recipes", RecipeHandler)
	http.HandleFunc("/api/recipe/", RecipeDetailHandler)
	http.HandleFunc("/recipe/", RecipeDetailPageHandler) // New route for serving the detailed HTML page
	http.HandleFunc("/api/add-favorite", AddFavoriteHandler)
}

// HomePageHandler serves the main HTML page when users visit the root URL.
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session-name")
	userID, ok := session.Values["userID"].(int)
	if !ok || userID == 0 {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, map[string]string{"ErrorMessage": "You must be logged in to view the profile page."})
		return
	}

	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session-name")
	userID, ok := session.Values["userID"].(int)
	if !ok || userID == 0 {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}

	userProfile, err := db.GetUserProfile(userID)
	if err != nil {
		http.Error(w, "Failed to load user profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProfile)
}

func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching random recipes...")
	recipes, err := api.GetRandomRecipes(5)
	if err != nil {
		http.Error(w, "Failed to get random recipes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

// RecipeDetailPageHandler serves the recipe detail HTML page.
func RecipeDetailPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/recipe_detail.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	recipeID := r.URL.Path[len("/recipe/"):] // Extract the recipe ID from the URL.
	tmpl.Execute(w, map[string]string{"RecipeID": recipeID})
}

func RecipeDetailHandler(w http.ResponseWriter, r *http.Request) {
	recipeID := r.URL.Path[len("/api/recipe/"):] // Extract the recipe ID from the URL.
	recipe, err := api.GetRecipeByID(recipeID)
	if err != nil {
		http.Error(w, "Failed to get recipe details", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	session, _ := session.Store.Get(r, "session-name")
	userID, ok := session.Values["userID"].(int)
	if !ok || userID == 0 {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}

	var req struct {
		RecipeID int    `json:"recipe_id"`
		Title    string `json:"title"`
		Image    string `json:"image"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := db.AddRecipeToFavorites(userID, req.RecipeID, req.Title, req.Image)
	if err != nil {
		http.Error(w, "Failed to add recipe to favorites", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Recipe added to favorites successfully"})
}
