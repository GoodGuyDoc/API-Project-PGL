package test

import (
	"spoonacular-api/api"
	"testing"
)

// All of the test for the API will be exported here so they can be run at once. -Christian Tester Taylor
func TestAllApi(t *testing.T) {
	api.TestGetRecipeByID(t)
	api.TestRandomRecipeCall(t)
	api.TestSimilarRecipe(t)
	api.TestLogToFile(t)
}
