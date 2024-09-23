package api_retrieval

import (
	// "errors"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
		panic(err) // Unsure what error actually looks like if there is one
	}
	// fmt.Println(resp.Body)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// body is byte array

	var jsonRes map[string]interface{} // declaring a map for key names as string and values as interface
	_ = json.Unmarshal(body, &jsonRes) // Unmarshalling
	// why is this set to _ above?

	// title := jsonRes["title"].(string) // convert value to string, we expect string

	// fmt.Println(title)

	// TODO: Fix this pls
	// recipes := jsonRes["recipes"].([]interface{})
	// recipes2 := recipes[0].(map[string]interface{})
	// recipes3 := recipes2[""]
	fmt.Printf("%s", jsonRes)

	// var j interface{}
	// err = json.NewDecoder(resp.Body).Decode(&j)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(j["Title"])
	// fmt.Printf("%s", j)

	close(ch)
}
