package api_retrieval

import (
	// "errors"

	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
	// "os"
	//"time"
)

const API_KEY string = "1eb4022c30c9430b8050b301be5339b9"

// Gets a random recipe based on a specific diet, meal type, or cuisine. These are provided with a slice of tags.
func GetRandomRecipe(tags []string) {
	result := make(chan string)
	go callRandomRecipe(tags, result)
	for i := range result {
		fmt.Println(i)
	}
}

// Internal function to call the spoontacular API random recipe with the provided param(s).
func callRandomRecipe(includes []string, ch chan string) {
	includetags := strings.Join(includes, ",")
	resp, err := http.Get("https://api.spoonacular.com/recipes/random?apiKey=" + API_KEY + "&number=1&include-tags=" + includetags)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// body is byte array
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := gjson.GetBytes(body, "recipes.0.vegetarian")

	fmt.Println(result)

	close(ch)
}
