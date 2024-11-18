package api

import (
	"errors"
	"fmt"
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

func GetSimilarRecipe(recipeId string, count int) ([]SimilarRecipe, error) {
	var recipes []SimilarRecipe

	for i := 0; i < 3; i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/%s/similar?apiKey=%s&number=%d", recipeId, API_KEY[i], count)
		err := sendApiCall(apiUrl, &recipes)

		if err != nil && err.Error() == "this api key is ratelimited" { // bad api key, go next
			continue
		} else if err != nil {
			return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
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
