package api

import (
	"fmt"
	"testing"
)

// fetches detailed information about a recipe using its ID
func GetRecipeByID(recipeID string) (*Recipe, error) {
	apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%s/information?apiKey=%s", recipeID, API_KEY)
	recipe, err := getRecipe(apiUrl)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func TestGetRecipeByID(t *testing.T){
	var err error = nil
	_, err = GetRecipeByID("1003464") //Makes a call to the API for testing purposes. Much easier than making entire set of mock data.
	if err != nil {
		t.Errorf("There was an error getting recipe by ID %v", err)
	}
	t.Log("Test GetRecipeByID Successful")
}
