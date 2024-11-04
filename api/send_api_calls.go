package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Takes an apiString input to call, then returns a *RecipeResponse, or an error
func getRecipeResponse(apiString string) (*RecipeResponse, error) {
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

// Takes an apiString input to call, then returns a *Recipe, or an error
func getRecipe(apiString string) (*Recipe, error) {
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

	var recipe Recipe
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &recipe, nil
}
