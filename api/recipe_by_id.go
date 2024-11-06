package api

import (
	"fmt"
)

// fetches detailed information about a recipe using its ID
func GetRecipeByID(recipeID string) (*Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%s/information?apiKey=%s", recipeID, API_KEY)
	recipe, err := getRecipe(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipe, nil
}
