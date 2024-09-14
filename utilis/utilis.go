package utilis

import (
	"log"
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
// func InsertUser(email, username, passwordHash string) error {
// 	_, err := DB.Exec(`INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)`, email, username, passwordHash)
// 	return err
// }

// HashPassword hashes the password using bcrypt.
