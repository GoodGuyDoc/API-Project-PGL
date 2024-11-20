package api

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

type SimilarRecipe struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	ImageType      string `json:"imageType"`
	ReadyInMinutes int    `json:"readyInMinutes"`
	Servings       int    `json:"servings"`
	SourceUrl      string `json:"sourceUrl"`
}

func GetSimilarRecipe(recipeId string, count int) ([]Recipe, error) {
	var similarRecipes []SimilarRecipe
	var recipes []Recipe

	for i := 0; i < 3; i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%s/similar?apiKey=%s&number=%d", recipeId, API_KEY[i], count)
		err := sendApiCall(apiUrl, &similarRecipes)

		if err != nil && err.Error() == "this api key is ratelimited" { // bad api key, go next
			continue
		} else if err != nil {
			return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
		}

		// make more queries to replace all similar recipes found with regular recipe objects
		for j := 0; j < len(similarRecipes); j++ {
			currRecipeId := strconv.Itoa(similarRecipes[j].ID) // if this conversion fails, we have problems
			currRecipe, err := GetRecipeByID(currRecipeId)
			if err != nil {
				return nil, fmt.Errorf("error making request to Spoonacular API: %v", err)
			}

			recipes = append(recipes, *currRecipe)
		}
		return recipes, nil
	}
	// if we did not find a good api key, throw an error (we finished looping)
	return nil, fmt.Errorf("error making request to Spoonacular API: %w", errors.New("all api keys are ratelimited"))
}

func TestSimilarRecipe(t *testing.T) error {
	//Random Similar Recipe call just for testing
	_, err := GetSimilarRecipe("122", 1)
	if err != nil {
		t.Errorf("There was an error grabbing similar recipe %v", err)
		return err
	}
	return nil
}
