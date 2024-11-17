package api

import (
	"fmt"
)

// fetches detailed information about a recipe using its ID
func GetRecipeByID(recipeID string) (*Recipe, error) {
	var recipe *Recipe
	var err error

	for i := 0; i < 3; i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%s/information?apiKey=%s", recipeID, API_KEY[i])
		recipe, err = getRecipe(apiUrl)

		if err.Error() == "this api key is ratelimited" {
			continue
		} else {
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipe, nil
}
