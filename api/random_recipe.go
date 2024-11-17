package api

import (
	"fmt"
	"os"
	"strings"
)

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
	var recipeResponse *RecipeResponse
	var err error

	// there are no while loops in go, lol
	for i := 0; i < 3; i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/random?apiKey=%s&number=%d", API_KEY[i], count)
		recipeResponse, err = getRecipeResponse(apiUrl)

		if err != nil && err.Error() == "this api key is ratelimited" {
			continue
		} else {
			break
		}
	}
	fmt.Printf("success")

	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipeResponse.Recipes, nil
}

// Returns count amount of random recipes from spoonacular api that are tagged with the specified tags
func GetRandomRecipesByTag(count int, tags []string) ([]Recipe, error) {
	includeTags := strings.Join(tags, ",")
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/random?apiKey=%s&number=%d&include-tags=%s", API_KEY[0], count, includeTags)
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
