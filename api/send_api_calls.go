package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

var API_KEY = [3]string{"a867e9b240a645c3a08192f8d6b8b61c", "7e40fb0f0f254a1aa1444150b5c71d07", "eea7c0c25e204d42b58aa324a6ddec5c"}

// Takes an apiString input to call, then returns a *RecipeResponse, or an error
func getRecipeResponse(apiString string) (*RecipeResponse, error) {
	resp, err := http.Get(apiString)
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
	err = logToFile("response_log.txt", indentedJSON) //TODO fix this with the new log format
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

// Sends an api call to spoonacular api. Provide the req string, as well as the address to a var of a struct type to be returned. Returns whether there was an error or not.
func sendApiCall(apiString string, returnVar interface{}) error {
	// Make sure we receive a pointer address
	if reflect.ValueOf(returnVar).Kind() != reflect.Ptr {
		return fmt.Errorf("returnVar must be a pointer")
	}

	resp, err := http.Get(apiString)
	// check for err
	if err != nil {
		return fmt.Errorf("error making request to Spoonacular API: %w", err)
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// convert body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, &returnVar)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	return nil
}
