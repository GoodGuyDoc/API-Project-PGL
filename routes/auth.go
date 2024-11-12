package routes

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"spoonacular-api/db"
	"spoonacular-api/session"

	"golang.org/x/crypto/bcrypt"
)

// routes for authentication-related actions
func SetupAuthRoutes() {
	// handle and serve static HTML pages
	http.HandleFunc("/register", RegisterPageHandler)
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/random_recipe", RandomRecipePageHandler)
	//handle and serve JSON data
	http.HandleFunc("/api/register", AddUserHandler)
	http.HandleFunc("/api/login", LoginHandler)
	http.HandleFunc("/api/logout", LogoutHandler)

}

// serves the login page for GET requests
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

// serves the registration page for GET requests
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/register.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

// handle user registration with a JSON-based API endpoint
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// define the request structure
	var req struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		Password  string `json:"password"`
	}
	// decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Hash password using bcrypt.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// store user data in db
	err = db.AddUser(req.Username, req.FirstName, string(hashedPassword))
	if err != nil {
		http.Error(w, "Error saving user to database", http.StatusInternalServerError)
		return
	}
	// send a JSON response back to the frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// handles user login with a JSON-based API
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Attempting login for user: %s", req.Username)

	// Retrieve hashed password from db
	hashedPassword, err := db.GetUserPassword(req.Username)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}

	// compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	log.Printf("User %s logged in successfully.", req.Username)

	// retrieve user profile to get ID
	userProfile, err := db.GetUserProfileByUsername(req.Username)
	if err != nil {
		http.Error(w, "Error retrieving user profile", http.StatusInternalServerError)
		return
	}

	// create a session and store user ID(used so users can only access their own data)
	session, _ := session.Store.Get(r, "session-name")
	session.Values["userID"] = userProfile.ID

	// Save the session.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

// handle user logout, clear the session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// clear the session
	session, _ := session.Store.Get(r, "session-name")
	session.Values["userID"] = nil
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

//handle routing for the random_recipe page routing
func RandomRecipePageHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        tmpl, err := template.ParseFiles("templates/random_recipe.html", "templates/header.html", "templates/footer.html")
        if err != nil {
            http.Error(w, "Error loading template", http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
        return
    }

    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
