package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"spoonacular-api/api"
	"strconv"
	"strings"
)

func SetupRecipeRoutes() {
	http.HandleFunc("/api/recipes", RecipeHandler)
	http.HandleFunc("/api/recipes/byTag", RecipeByTagHandler)
	http.HandleFunc("/api/recipe/", RecipeDetailHandler)
	http.HandleFunc("/recipe_detail", RecipeDetailPageHandler)
	http.HandleFunc("/api/convert", ConversionHandler)
	http.HandleFunc("/random_recipe_page", RandomRecipePageHandler)
}

// RecipeHandler fetches random recipes and serves them as JSON back to the frontend.
func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching random recipes...")
	recipes, err := api.GetRandomRecipes(5) //get 5 random recipes
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Failed to get random recipes"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func RecipeByTagHandler(w http.ResponseWriter, r *http.Request) {
	// parse count
	countParam := r.URL.Query().Get("count")
	count := 5 // Default to 5 if no count given
	if countParam != "" {
		parsedCount, err := strconv.Atoi(countParam)
		if err == nil {
			count = parsedCount
		}
	}

	// parse tags
	tagsParam := r.URL.Query().Get("tags")
	tags := []string{}
	if tagsParam != "" {
		tags = strings.Split(tagsParam, ",")
	}

	recipes, err := api.GetRandomRecipesByTag(count, tags)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Failed to get recipes by tag"}`, http.StatusInternalServerError)
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

func ConversionHandler(w http.ResponseWriter, r *http.Request) {
	ingredientName := r.URL.Query().Get("ingredientName")
	amountParam := r.URL.Query().Get("amount")
	unit := r.URL.Query().Get("unit")
	convertToUnit := r.URL.Query().Get("convertToUnit")

	if ingredientName == "" || amountParam == "" || unit == "" || convertToUnit == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	//convert amount to float
	amount, err := strconv.ParseFloat(amountParam, 64)
	if err != nil {
		http.Error(w, "Invalid amount parameter", http.StatusBadRequest)
		return
	}

	conversionInfo, err := api.ConvertAmount(ingredientName, amount, unit, convertToUnit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to convert amount: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conversionInfo)
}
func RandomRecipePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/random_recipe_page.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Error loading template"}`, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
