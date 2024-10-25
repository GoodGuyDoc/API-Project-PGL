package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// fetches detailed information about a recipe using its ID
func GetRecipeByID(recipeID string) (*Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%s/information?apiKey=%s", recipeID, API_KEY)
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

	var recipe Recipe
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &recipe, nil
}
