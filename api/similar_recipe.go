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

TestSimilarRecipe(t *testing.T) {
	//Random Similar Recipe call just for testing
	var err := nil
	_, err := GetSimilarRecipe("122",4)
	if err != nil{
		err = fmt.Errorf("error getting Similar Recipe mid num API: %v", err)
}
	_, err := GetSimilarRecipe("4001",1)
if err != nil{
err = fmt.Errorf("error getting Similar Recipe 1 num API: %v", err)
}
	_, err := GetSimilarRecipe("4001",10)
if err != nil{
err = fmt.Errorf("error getting Similar Recipe large num API: %v", err)
}
	if err != nil{
		t.Errorf("error getting spoonacular recipe with id 122: %v", err)
	}
}

