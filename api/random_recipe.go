package api

import (
	"fmt"
	"os"
	"strings"
)

const API_KEY = "a867e9b240a645c3a08192f8d6b8b61c"

type RecipeResponse struct {
	Recipes []Recipe `json:"recipes"`
}

type Recipe struct {
	ID                   int                   `json:"id"`
	Title                string                `json:"title"`
	Image                string                `json:"image"`
	AnalyzedInstructions []AnalyzedInstruction `json:"analyzedInstructions"`
	ExtendedIngredients  []Ingredient          `json:"extendedIngredients"`
}

type Ingredient struct {
	Original string `json:"original"`
}

type AnalyzedInstruction struct {
	Name  string `json:"name"`
	Steps []Step `json:"steps"`
}

type Step struct {
	Number int    `json:"number"`
	Step   string `json:"step"`
}

// Returns count amount of random recipes from spoonacular api
func GetRandomRecipes(count int) ([]Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/random?apiKey=%s&number=%d", API_KEY, count)
	recipeResponse, err := getRecipeResponse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipeResponse.Recipes, nil
}

// Returns count amount of random recipes from spoonacular api that are tagged with the specified tags
func GetRandomRecipesByTag(count int, tags []string) ([]Recipe, error) {
	includeTags := strings.Join(tags, ",")
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/random?apiKey=%s&number=%d&include-tags=%s", API_KEY, count, includeTags)
	recipeResponse, err := getRecipeResponse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipeResponse.Recipes, nil
}

// logToFile writes data to a file.
func logToFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing data to file: %w", err)
	}

	// Write a newline for better readability between entries
	_, err = file.WriteString("\n\n")
	if err != nil {
		return fmt.Errorf("error writing newline to file: %w", err)
	}

	return nil
}
