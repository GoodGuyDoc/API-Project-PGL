package main

import (
	api "API-Project-PGL/api_retrieval"
	//"API-Project-PGL/fileops"
	"fmt"
	//"strings"
	"html/template"
	"log"
	"net/http"
)

func main() {
	/* commenting out code w/ command line aspects for http testing
	for {
		menu()
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		err = fileops.WriteToFile("testfile.txt", "Hello, World!")
		if err != nil {
			fmt.Println(err)
		}

		switch choice {
		case 1:
			handleRandomRecipe()
		case 2:
			handleFindRecipesByIngredients()
		case 3:
			handleFindRecipesByNutrients()
		case 4:
			handleAddToMealPlan()
		case 5:
			handleDeleteFromMealPlan()
		case 6:
			fmt.Println("Exiting the application.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
	*/

	/*
		This is very rough code
		Just showing some stuff w/ http
		Everything can be changed and probably
		should be moved to other places :]
	*/

	tags := map[string][]string{
		"Tags": {"dessert", "italian"},
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, tags)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		var responseTags []string

		for key := range tags["Tags"] {

			if r.PostFormValue(tags["Tags"][key]) == "on" {
				responseTags = append(responseTags, tags["Tags"][key])
			}

		}
		handleRandomRecipe(responseTags)
		fmt.Println()
		fmt.Println()
		// the template isn't set up yet
		// Functions wil have to be edited to return the data
		// tmpl.Execute(w, whateverAPIwillReturn)
		// for now, just redirect back to the form

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/send-tag/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func menu() {
	fmt.Println("=== Meal Planner Menu ===")
	fmt.Println("1. Get a random recipe")
	fmt.Println("2. Find recipes by ingredients")
	fmt.Println("3. Find recipes by nutrients")
	fmt.Println("4. Add an item to meal plan")
	fmt.Println("5. Delete an item from meal plan")
	fmt.Println("6. Exit")
}

func handleRandomRecipe(tags []string) {
	api.GetRandomRecipe(tags)
}

/* commenting out code w/ errors for http testing

func handleFindRecipesByIngredients() {
	var input string
	fmt.Print("Enter ingredients separated by commas (e.g., apples, flour, sugar): ")
	fmt.Scan(&input)
	ingredients := strings.Split(input, ",")
	for i := range ingredients {
		ingredients[i] = strings.TrimSpace(ingredients[i]) // Trim spaces
	}

	recipes, err := api.FindRecipesByIngredients(ingredients, 2)
	if err != nil {
		fmt.Println("Error finding recipes by ingredients:", err)
		return
	}

	fmt.Println("Recipes found by ingredients:")
	for _, recipe := range recipes {
		fmt.Printf("Title: %s, Likes: %d, Image: %s\n", recipe.Title, recipe.Likes, recipe.Image)
		for _, missed := range recipe.MissedIngredients {
			fmt.Printf("Missed Ingredient: %s (%s)\n", missed.Name, missed.Original)
		}
		for _, used := range recipe.UsedIngredients {
			fmt.Printf("Used Ingredient: %s (%s)\n", used.Name, used.Original)
		}
		fmt.Println()
	}
}

func handleFindRecipesByNutrients() {
	var minCarbs, maxCarbs, number int
	fmt.Print("Enter minimum carbs: ")
	fmt.Scan(&minCarbs)
	fmt.Print("Enter maximum carbs: ")
	fmt.Scan(&maxCarbs)
	fmt.Print("Enter number of recipes to find: ")
	fmt.Scan(&number)

	recipes, err := api.FindRecipesByNutrients(minCarbs, maxCarbs, number)
	if err != nil {
		fmt.Println("Error finding recipes by nutrients:", err)
		return
	}

	fmt.Println("Recipes found by nutrients:")
	for _, recipe := range recipes {
		fmt.Printf("Title: %s, Calories: %d, Carbs: %s, Image: %s\n", recipe.Title, recipe.Calories, recipe.Carbs, recipe.Image)
	}
}

func handleAddToMealPlan() {
	var username, hash, itemName string
	var date, slot, position int

	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	fmt.Print("Enter your hash: ")
	fmt.Scan(&hash)

	fmt.Print("Enter the item name to add (e.g., '1 banana'): ")
	fmt.Scan(&itemName)

	fmt.Print("Enter the date (timestamp): ")
	fmt.Scan(&date)
	fmt.Print("Enter the slot (1 for breakfast, 2 for lunch, 3 for dinner): ")
	fmt.Scan(&slot)
	fmt.Print("Enter the position (0 for top): ")
	fmt.Scan(&position)

	// Create the meal plan item
	item := api.MealPlanItem{
		Date:     int64(date),
		Slot:     slot,
		Position: position,
		Type:     "INGREDIENTS",
		Value: api.ItemValue{
			Ingredients: []api.Ingredient{{Name: itemName}},
		},
	}

	// Call the API function to add the item to the meal plan
	err := api.AddMealPlanItem(username, hash, item)
	if err != nil {
		fmt.Println("Error adding item to meal plan:", err)
		return
	}

	fmt.Println("Item added to meal plan successfully.")
}

func handleDeleteFromMealPlan() {
	var username, hash string
	var itemID int

	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	fmt.Print("Enter your hash: ")
	fmt.Scan(&hash)
	fmt.Print("Enter the item ID to delete: ")
	fmt.Scan(&itemID)

	// Call the API function to delete the item from the meal plan
	err := api.DeleteMealPlanItem(username, itemID, hash)
	if err != nil {
		fmt.Println("Error deleting item from meal plan:", err)
		return
	}

	fmt.Println("Item deleted from meal plan successfully.")
}
*/
