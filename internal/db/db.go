package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"golangApp-postgres/internal/models"
)

var DB *sql.DB

// InitDB connects to PostgreSQL using environment variables
func InitDB() error {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)

	log.Printf("📡 Connecting to DB host: %s, DB name: %s", os.Getenv("DBHOST"), os.Getenv("DBNAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("sql open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("db ping error: %w", err)
	}

	DB = db
	log.Println("📦 Connected to PostgreSQL")
	return nil
}

// GetAllUsers fetches all users
func GetAllUsers() ([]models.User, error) {
	rows, err := DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Printf("Scan failed: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID fetches a user by ID
func GetUserByID(id int) (models.User, error) {
	var u models.User
	err := DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

// InsertUser creates a new user
func InsertUser(name, email string) error {
	_, err := DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	return err
}

// UpdateUser updates name and email by ID
func UpdateUser(id int, name, email string) error {
	_, err := DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", name, email, id)
	return err
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
