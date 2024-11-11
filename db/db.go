package db

import (
	"database/sql"
	"fmt"
	"time"

	//_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
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
func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite", "./db/gogrub.db")
	if err != nil {
		return err
	}

	sqlStatement := `CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		first_name VARCHAR(50) NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS Recipes (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		recipe_id INTEGER NOT NULL,
		title VARCHAR(255) NOT NULL,
		image TEXT NOT NULL,
		date_added TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS Tags (
		id INTEGER PRIMARY KEY,
		tag VARCHAR(255) UNIQUE
	);

	CREATE TABLE IF NOT EXISTS UserTags (
		user_id INTEGER NOT NULL,
		tag_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
		FOREIGN KEY (tag_id) REFERENCES Tags(id) ON DELETE CASCADE,
		PRIMARY KEY(user_id, tag_id)
	);

	INSERT OR IGNORE INTO Tags (tag) 
	VALUES ('breakfast'), ('lunch'), ('dinner');
	`
	DB.Exec(sqlStatement)

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
        VALUES ($1, $2, $3, $4, datetime('now'))
    `
	_, err := DB.Exec(query, userID, recipeID, title, image)
	if err != nil {
		return fmt.Errorf("error adding recipe to favorites: %w", err)
	}
	return nil
}

func AddTagToFavorites(userID, tagID int) error {
	query := `
        INSERT INTO UserTags (user_id, tag_id)
        VALUES ($1, $2)
    `
	_, err := DB.Exec(query, userID, tagID)
	if err != nil {
		return fmt.Errorf("error adding tag to favorites: %w", err)
	}
	return nil
}

func RemoveTagFromFavorites(userID, tagID int) error {
	query := `
        DELETE FROM UserTags 
		WHERE user_id = $1 AND tag_id = $1;
    `
	_, err := DB.Exec(query, userID, tagID)
	if err != nil {
		return fmt.Errorf("error removing tag from favorites: %w", err)
	}
	return nil
}

func ClearFavoriteTags(userID int) error {
	query := `
		DELETE FROM UserTags 
		WHERE user_id = $1;
	`
	_, err := DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error clearing favorite tags: %w", err)
	}
	return nil
}

func GetUserTags(userID int) ([]string, error) {

	var userTags []string

	rows, err := DB.Query(`
        SELECT tag FROM UserTags 
		JOIN Tags ON tag_id = id
        WHERE user_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching favorite tags: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, fmt.Errorf("error scanning favorite tag: %w", err)
		}
		userTags = append(userTags, tag)
	}

	return userTags, nil
}

func GetAllTags() ([]string, error) {
	var tags []string

	rows, err := DB.Query(`
        SELECT tag FROM Tags`)
	if err != nil {
		return nil, fmt.Errorf("error fetching tags: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, fmt.Errorf("error scanning tag: %w", err)
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
