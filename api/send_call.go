package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Takes an apiString, sends the call, and processes the response. Provides a pointer to the RecipeResponse.
func send_api_call(apiString string) (*RecipeResponse, error) {
	resp, err := http.Get(apiString)
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var prettyJSON map[string]interface{}
	err = json.Unmarshal(body, &prettyJSON)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON for pretty print: %w", err)
	}

	indentedJSON, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error formatting JSON: %w", err)
	}

	// Log the formatted JSON to a file named response_log.txt
	err = logToFile("response_log.txt", indentedJSON)
	if err != nil {
		return nil, fmt.Errorf("error logging to file: %w", err)
	}

	var recipeResponse RecipeResponse
	err = json.Unmarshal(body, &recipeResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	if len(recipeResponse.Recipes) == 0 {
		return nil, fmt.Errorf("no recipes found")
	}

	return &recipeResponse, nil
}
