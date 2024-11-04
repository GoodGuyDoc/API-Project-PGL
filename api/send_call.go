package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// Takes an apiString, sends the call, and processes the response. Provides a pointer to the RecipeResponse.
func send_api_call(apiString string, expectMoreThanOneRecipe bool) (*RecipeResponse, error) {
	if expectMoreThanOneRecipe {
		return send_multiple(apiString)
	} else {
		recipe, err := send_single(apiString)
		recipeResponse := &RecipeResponse{
			Recipes: []Recipe{},
		}
		recipeResponse.Recipes = append(recipeResponse.Recipes, *recipe)
		return recipeResponse, err
	}
}

func send_multiple(apiString string) (*RecipeResponse, error) {
	resp, err := http.Get(apiString)
	if err != nil {
		return nil, fmt.Sprintf("error making request to Spoonacular API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Sprintf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Sprintf("error reading response body: %v", err)
	}

	var prettyJSON map[string]interface{}
	err = json.Unmarshal(body, &prettyJSON)
	if err != nil {
		return nil, fmt.Sprintf("error unmarshaling JSON for pretty print: %v", err)
	}

	indentedJSON, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		return nil, fmt.Sprintf("error formatting JSON: %v", err)
	}

	// Log the formatted JSON to a file named response_log.txt
	err = logToFile("response_log.txt", indentedJSON)
	if err != nil {
		return nil, fmt.Sprintf("error logging to file: %v", err)
	}

	var recipeResponse RecipeResponse
	err = json.Unmarshal(body, &recipeResponse)
	if err != nil {
		return nil, fmt.Sprintf("error parsing JSON: %v", err)
	}

	if len(recipeResponse.Recipes) == 0 {
		return nil, "no recipes found"
	}

	return &recipeResponse, ""
}

func TestApiCall(apiString string,*t testing.T){
	var retStr := ""
	res,retStr := send_api_call(apiString)
	if retStr != "" {
		t.Fatalf(retStr)
	}
}
func send_single(apiString string) (*Recipe, error) {
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
