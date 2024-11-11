package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// Takes an apiString, sends the call, and processes the response. Returns RecipeResponse pointer and error string.
func send_api_call(apiString string) (*RecipeResponse, string) {
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
