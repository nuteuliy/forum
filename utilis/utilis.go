package utilis

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// UserExists checks if the user with the given email already exists.
func UserExists(email string) bool {
	var count int

	row := DB.QueryRow(`SELECT COUNT(*) FROM users WHERE email = ?`, email)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}
func UserNameExist(username string) bool {
	var count int
	row := DB.QueryRow(`SELECT COUNT(*) FROM users WHERE username = ?`, username)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

// InsertUser inserts a new user into the database.
func InsertUser(email, username, passwordHash string) error {
	_, err := DB.Exec(`INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)`, email, username, passwordHash)
	return err
}

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func CompareHashedPassword(username string, hashedPassword *string) error {

	err := DB.QueryRow(`SELECT password_hash FROM users WHERE username = ?`, username).Scan(hashedPassword)

	if err == sql.ErrNoRows {
		return err
	} else {
		return nil
	}
}
