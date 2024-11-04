package api

import (
	"fmt"
)

// fetches detailed information about a recipe using its ID
func GetRecipeByID(recipeID string) (*Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%s/information?apiKey=%s", recipeID, API_KEY)
	recipeResponse, err := send_api_call(apiUrl, false)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return &recipeResponse.Recipes[0], nil
}
