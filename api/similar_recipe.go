package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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
		resp, err := http.Get(apiUrl)

		if err != nil && err.Error() == "this api key is ratelimited" { // bad api key, go next
			continue
		} else {
			if err != nil {
				return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode == 402 || resp.StatusCode == 429 {
					return nil, fmt.Errorf("this api key is ratelimited")
				}
				return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			err = json.Unmarshal(body, &recipes)
			if err != nil {
				return nil, fmt.Errorf("error parsing JSON: %w", err)
			}
			return recipes, nil
		}
	}
	// if we did not find a good api key, throw an error (we finished looping)
	return nil, fmt.Errorf("error making request to Spoonacular API: %w", errors.New("all api keys are ratelimited"))
}
