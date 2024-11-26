package api

import (
	"fmt"
)

// fetches detailed information about a recipe using its ID
func GetRecipeByID(recipeID string) (*Recipe, error) {
	var recipe *Recipe
	var err error

	curr_api_key := getAPIKeys() // make sure keys are initialized

	for i := 0; i < len(curr_api_key); i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%s/information?apiKey=%s", recipeID, curr_api_key[i])
		recipe, err = getRecipe(apiUrl)

		if err != nil {
			if err.Error() == "this api key is ratelimited" {
				continue
			} else if err.Error() == "no recipe found" {
				return &Recipe{}, nil // return empty struct
			} else {
				return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
			}
		}
		break
	}

	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipe, nil
}
