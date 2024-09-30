package api_retrieval

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

const API_KEY string = "1eb4022c30c9430b8050b301be5339b9"

// TODO: Fill this entirely or as we need more data
// Struct defining all of the important data we care about from a recipe
type Recipe struct {
	vegetarian     bool
	vegan          bool
	veryHealthy    bool
	veryPopular    bool
	readyInMinutes int
	servings       int
	title          string
	summary        string
	instructions   string
}

// Gets a random recipe based on a specific diet, meal type, or cuisine. These are provided with a slice of tags.
func GetRandomRecipe(tags []string) {
	recipeChannel := make(chan *Recipe)
	go callRandomRecipe(tags, recipeChannel)
	for recipe := range recipeChannel {
		// TODO: Remove this eventually
		fmt.Printf("i points to: %p, i's address is: %p\n", recipe, &recipe)
		// TODO: Send the data to another function to print
		fmt.Printf("The title of the recipe you got: %s", recipe.title)
	}
}

// Gets a random random recipe with the provided param(s) and sends the data through the string channel
func callRandomRecipe(includes []string, ch chan *Recipe) {
	includetags := strings.Join(includes, ",")
	jsonData, err := getSpoonData("https://api.spoonacular.com/recipes/random?apiKey=" + API_KEY + "&number=1&include-tags=" + includetags)
	if err != nil {
		panic(err)
	}

	// TODO: Is there a better way to iterate through each part of the struct and grab the value?
	recipe := Recipe{}
	val := gjson.GetBytes(jsonData, "recipes.0.vegetarian")
	recipe.vegetarian = val.Bool()
	recipe.vegan = gjson.GetBytes(jsonData, "recipes.0.vegan").Bool()
	recipe.veryHealthy = gjson.GetBytes(jsonData, "recipes.0.veryHealthy").Bool()
	recipe.veryPopular = gjson.GetBytes(jsonData, "recipes.0.veryPopular").Bool()
	recipe.readyInMinutes = int(gjson.GetBytes(jsonData, "recipes.0.readyInMinutes").Int())
	recipe.servings = int(gjson.GetBytes(jsonData, "recipes.0.servings").Int())
	recipe.title = gjson.GetBytes(jsonData, "recipes.0.title").String()
	recipe.summary = gjson.GetBytes(jsonData, "recipes.0.summary").String()
	recipe.instructions = gjson.GetBytes(jsonData, "recipes.0.instructions").String()

	ch <- &recipe
	close(ch)
}
