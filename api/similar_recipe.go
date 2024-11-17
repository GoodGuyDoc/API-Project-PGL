package api

import (
	"fmt"
)

func GetSimilarRecipe(recipeId int) ([]Recipe, error) {
	var recipeResponse *RecipeResponse
	var err error

	for i := 0; i < 3; i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%d/similar?apiKey=%s", recipeId, API_KEY[i])
		recipeResponse, err = getRecipeResponse(apiUrl)

		if err.Error() == "this api key is ratelimited" {
			continue
		} else {
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipeResponse.Recipes, nil
}
