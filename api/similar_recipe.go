package api

import (
	"fmt"
	"testing"
)

func GetSimilarRecipe(recipeId int) ([]Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%d/similar?apiKey=%s", recipeId, API_KEY)
	recipeResponse, err := send_api_call(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return recipeResponse.Recipes, nil
}

TestSimilarRecipe(t *testing.T)(error,[]Recipe recipe){
	//Random Similar Recipe call just for testing
	var err := nil
	recipe,err := GetSimilarRecipe(122)
	if err != nil{
		t.Fatal()
	}
	return recipe,err
}

