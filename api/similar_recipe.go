package api

import (
	"fmt"
	"testing"
)

func GetSimilarRecipe(recipeId int) ([]Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%d/similar?apiKey=%s", recipeId, API_KEY)
	recipeResponse, err := getRecipeResponse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipeResponse.Recipes, nil
}

func TestSimilarRecipe(t *testing.T) error {
	//Random Similar Recipe call just for testing
	_, err := GetSimilarRecipe(122)
	if err != nil {
		t.Errorf("There was an error grabbing similar recipe %v", err)
		return err
	}
	return nil
}
