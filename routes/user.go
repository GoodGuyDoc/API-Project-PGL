package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"spoonacular-api/db"
	"spoonacular-api/session"
)

func SetupUserRoutes() {
	//handle and serve static HTML pages(accessable to user)
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/profile", ProfilePageHandler)

	//handle and serve JSON data(accessed programmatically within the html pages)
	http.HandleFunc("/api/profile", ProfileHandler)
	http.HandleFunc("/api/add-favorite", AddFavoriteHandler)
}

// HomePageHandler serves the main HTML page when users visit the root URL.
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// serves the static profile page
func ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session-name")
	userID, ok := session.Values["userID"].(int)

	//if user is not logged in, redirect to login page
	if !ok || userID == 0 {
		tmpl, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, map[string]string{"ErrorMessage": "You must be logged in to view the profile page."})
		return
	}

	tmpl, err := template.ParseFiles("templates/profile.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// ProfileHandler fetches the user's profile data and serves it as JSON back to the frontend.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session-name")
	userID, ok := session.Values["userID"].(int)
	if !ok || userID == 0 {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}

	userProfile, err := db.GetUserProfile(userID)
	if err != nil {
		http.Error(w, "Failed to load user profile", http.StatusInternalServerError)
		return
	}
	//send json response back to the frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProfile)
}

// AddFavoriteHandler handles adding a recipe to the user's favorites list.
func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	session, _ := session.Store.Get(r, "session-name")
	userID, ok := session.Values["userID"].(int)
	if !ok || userID == 0 {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}

	var req struct {
		RecipeID int    `json:"recipe_id"`
		Title    string `json:"title"`
		Image    string `json:"image"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//add recipe to favorites table in the database(linked to user ID)
	err := db.AddRecipeToFavorites(userID, req.RecipeID, req.Title, req.Image)
	if err != nil {
		http.Error(w, "Failed to add recipe to favorites", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Recipe added to favorites successfully"})
}
