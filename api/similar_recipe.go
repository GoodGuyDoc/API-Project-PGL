package api

import (
	"fmt"
)

func GetSimilarRecipe(recipeId int) ([]Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%d/similar?apiKey=%s", recipeId, API_KEY)
	recipeResponse, err := send_api_call(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipeResponse.Recipes, nil
}
