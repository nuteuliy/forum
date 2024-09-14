package utilis

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func CompareHashedPassword(username string, hashedPassword *string,userID *int) error {

	err := DB.QueryRow(`SELECT password_hash,id FROM users WHERE username = ?`, username).Scan(hashedPassword,userID)

	if err == sql.ErrNoRows {
		return err
	} else {
		return nil
	}
}
