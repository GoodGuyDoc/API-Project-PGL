package test

import (
	"spoonacular-api/api"
	"testing"
)

// All of the test for the API will be exported here so they can be run at once. -Christian Tester Taylor

func TestConvertAmount(t *testing.T) {
	var (
		ingredientName string  = "flour"
		sourceAmount   float64 = 2.5
		sourceUnit     string  = "cups"
		targetUnit     string  = "grams"
	)
	_, err := api.ConvertAmount(ingredientName, sourceAmount, sourceUnit, targetUnit)
	if err != nil {
		t.Errorf("There was an error in the converting testCase %v", err)
	}
	t.Log("Convert Amount Test Successful")
	_, err = api.ConvertAmount(ingredientName, sourceAmount, "ERRORAMOUNT!----", targetUnit)
	if err != nil {
		t.Errorf("Error should be thrown here as ERRORAMOUNT!----- is not a valid source unit for api call")
	} else {
		t.Log("Test case for incorrect string successful")
	}
	t.Log("Testing for Conversion Complete")
}

func TestRandomRecipeCall(t *testing.T) {
	_, err := api.GetRandomRecipes(100)
	if err != nil {
		t.Errorf("There was an error in random recipe testing ERROR: %v", err)
	} else {
		t.Log("RandomRecipe Valid Call Successful.")
	}
	_, err = api.GetRandomRecipes(-5)
	if err != nil {
		t.Errorf("Random Recipe call with negative number did not return an error.")
	} else {
		t.Logf("RandomRecipe call with negative number did return a number test successful ERROR:%v", err)
	}
	t.Log("Testing for RandomRecipes Complete...\n")
}

func TestGetRecipeByID(t *testing.T) {
	var err error = nil
	_, err = api.GetRecipeByID("1003464") //Makes a call to the API for testing purposes. Much easier than making entire set of mock data.
	if err != nil {
		t.Errorf("There was an error getting recipe by ID %v", err)
	}
	t.Log("Test GetRecipeByID Successful")
	_, err = api.GetRecipeByID("999999999999999")
	if err != nil {
		t.Logf("Max Number length created err successfully, ERROR %v", err)
	} else {
		t.Error("Recipe ID of this length should return an error")
	}
	t.Log("Test RandomRecipeByID COMPLETED...\n")

}
func TestSimilarRecipe(t *testing.T) {
	// emptRecipe := api.Recipe{}
	//Random Similar Recipe call just for testing
	_, err := api.GetSimilarRecipe("122", 1)
	if err != nil {
		t.Errorf("There was an error grabbing similar recipe %v", err)
	} else {
		t.Log("Basic similar recipe test complete")
	}
	// _, err = api.GetSimilarRecipe("100", -5)
	// if err != nil {
	// 	t.Logf("Test Similar Recipe with negative number returned error or empty struct correctly ERROR:%v", err)
	// } else {
	// 	t.Error("Test Similar Recipe with a negative number should have returned an error but did not")
	// }
	t.Log("Testing finished for Similar Recipe")
}

func TestGetRandomRecipeByTag(t *testing.T) {
	// emptRecipe := api.Recipe{}
	//Testing with proper tags
	_, err := api.GetRandomRecipesByTag(1, "vegetarian", "dairy")
	if err != nil {
		t.Errorf("There was an error when putting in the template recipeTag information ERROR:%v", err)
	} else {
		t.Log("Test getRandomRecipeByTag Successful")
	}
	//Test with improper tags
	// responseStruct, err := api.GetRandomRecipesByTag(1, "BUTTERMYBISCUT", "LoveIsland")
	// if err != nil || reflect.DeepEqual(emptRecipe, responseStruct) { //If there is an error or an empty struct
	// 	t.Log("Test of incorrect request successful")
	// } else {
	// 	t.Error("Test of incorrect request unsuccessful")
	// }
	// t.Log("Testing for GetRandomRecipeByTag Complete")
}

// func TestLogToFile(t *testing.T) {
// 	tempFile, err := os.CreateTemp("./test", "testlogfile*") //Create the temp file
// 	if err != nil {
// 		t.Errorf("There was an error when creating the testing tempFile: %v", err)
// 	}
// 	defer func() {
// 		tempFile.Close()
// 		os.Remove(tempFile.Name())
// 	}()
// 	fileAbs, err := filepath.Abs(tempFile.Name())
// 	err = api.LogToFile(fileAbs, []byte("Hello World!"))

// 	if err != nil {
// 		t.Errorf("There was an error while testing file logging %v", err)
// 	}
// }
