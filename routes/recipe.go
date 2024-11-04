package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"spoonacular-api/api"
)

func SetupRecipeRoutes() {
	http.HandleFunc("/api/recipes", RecipeHandler)
	http.HandleFunc("/api/recipe/", RecipeDetailHandler)
	http.HandleFunc("/recipe/", RecipeDetailPageHandler)
}

// RecipeHandler fetches random recipes and serves them as JSON back to the frontend.
func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching random recipes...")
	recipes, err := api.GetRandomRecipes(5) //get 5 random recipes
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Failed to get random recipes"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

// RecipeDetailPageHandler serves the recipe detail HTML page.
func RecipeDetailPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/recipe_detail.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Error loading template"}`, http.StatusInternalServerError)
		return
	}

	recipeID := r.URL.Path[len("/recipe/"):] // Extract the recipe ID from the URL.
	tmpl.Execute(w, map[string]string{"RecipeID": recipeID})
}

// RecipeDetailHandler fetches the details of a specific recipe and serves them as JSON back to the frontend.
func RecipeDetailHandler(w http.ResponseWriter, r *http.Request) {
	recipeID := r.URL.Path[len("/api/recipe/"):] //extract the recipe ID from the URL
	recipe, err := api.GetRecipeByID(recipeID)   //get recipe details by ID
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Failed to get recipe details"}`, http.StatusInternalServerError)
		return
	}

	//send json response back to the frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}
