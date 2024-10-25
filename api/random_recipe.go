package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

// fetches a random recipe from the Spoonacular API
func GetRandomRecipes(count int) ([]Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/random?apiKey=%s&number=%d", API_KEY, count)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Pretty-print the JSON response using json.MarshalIndent
	var prettyJSON map[string]interface{}
	err = json.Unmarshal(body, &prettyJSON)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON for pretty print: %w", err)
	}

	indentedJSON, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error formatting JSON: %w", err)
	}

	// Log the formatted JSON to a file named response_log.txt
	err = logToFile("response_log.txt", indentedJSON)
	if err != nil {
		return nil, fmt.Errorf("error logging to file: %w", err)
	}

	var recipeResponse RecipeResponse
	err = json.Unmarshal(body, &recipeResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	if len(recipeResponse.Recipes) == 0 {
		return nil, fmt.Errorf("no recipes found")
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
