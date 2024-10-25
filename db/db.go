package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type UserProfile struct {
	ID        int
	FirstName string
	Username  string
	Recipes   []FavoriteRecipe
}

type FavoriteRecipe struct {
	RecipeID int
	Title    string
	Image    string
	AddedAt  time.Time
}

var DB *sql.DB

// initialize db connections
func InitDB(host, port, user, password, dbname string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	return DB.Ping()
}

// insert new user into the Users table
func AddUser(username, firstName, passwordHash string) error {
	sqlStatement := `
        INSERT INTO Users (username, first_name, password_hash)
        VALUES ($1, $2, $3)`
	_, err := DB.Exec(sqlStatement, username, firstName, passwordHash)
	return err
}

// retrieves hashed password for a given username
func GetUserPassword(username string) (string, error) {
	var hashedPassword string
	sqlStatement := `SELECT password_hash FROM Users WHERE username = $1`
	err := DB.QueryRow(sqlStatement, username).Scan(&hashedPassword)
	return hashedPassword, err
}

// fetches the user's first name, username, and favorite recipes
func GetUserProfile(userID int) (*UserProfile, error) {
	userProfile := &UserProfile{}

	// Fetch user details
	err := DB.QueryRow("SELECT first_name, username FROM Users WHERE id = $1", userID).
		Scan(&userProfile.FirstName, &userProfile.Username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user details: %w", err)
	}

	// Fetch users favorite recipes
	rows, err := DB.Query(`
        SELECT recipe_id, title, image, date_added
        FROM Recipes
        WHERE user_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching favorite recipes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var recipe FavoriteRecipe
		if err := rows.Scan(&recipe.RecipeID, &recipe.Title, &recipe.Image, &recipe.AddedAt); err != nil {
			return nil, fmt.Errorf("error scanning favorite recipe: %w", err)
		}
		userProfile.Recipes = append(userProfile.Recipes, recipe)
	}

	return userProfile, nil
}

func GetUserProfileByUsername(username string) (*UserProfile, error) {
	userProfile := &UserProfile{}

	// Fetch user details including the ID.
	err := DB.QueryRow("SELECT id, first_name, username FROM Users WHERE username = $1", username).
		Scan(&userProfile.ID, &userProfile.FirstName, &userProfile.Username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user profile: %w", err)
	}

	return userProfile, nil
}

func AddRecipeToFavorites(userID, recipeID int, title, image string) error {
	query := `
        INSERT INTO Recipes (user_id, recipe_id, title, image, date_added)
        VALUES ($1, $2, $3, $4, NOW())
    `
	_, err := DB.Exec(query, userID, recipeID, title, image)
	if err != nil {
		return fmt.Errorf("error adding recipe to favorites: %w", err)
	}
	return nil
}
