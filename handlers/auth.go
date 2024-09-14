package handlers

import (
	"FORUM/utilis"
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		// fmt.Println(r.Header["Content-Type"])
		// email := r.FormValue("email")
		// username := r.FormValue("username")
		// password := r.FormValue("password")

		// if utilis.UserExists(email) {
		// 	http.Error(w, "Email already exists", http.StatusConflict)
		// 	return
		// }
		var user struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if utilis.UserExists(user.Email) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"emailError": "Email already exists.",
			})
			return

		}
		if utilis.UserNameExist(user.Username) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"usernameError": "Username already exists.",
			})
			return
		}
		hashedPassword, err := utilis.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "HashingPasswordError", http.StatusInternalServerError)

			return
		}

		// Insert the user into the database
		err = utilis.InsertUser(user.Email, user.Username, hashedPassword)
		if err != nil {
			http.Error(w, "Error inserting user", http.StatusInternalServerError)
			return
		}
		// v := json.NewDecoder(r.Body)
		// fmt.Println(v, "hello ")

		// _, err = utilis.DB.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)",
		// 	user.Email, user.Username, hashedPassword)
		// if err != nil {
		// 	http.Error(w, "Error inserting user", http.StatusInternalServerError)
		// 	return
		// }

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		var userLogin struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
			http.Error(w, "HashingPasswordError", http.StatusInternalServerError)
			return
		}

		var hashedPassword string
		var userID int
		// if !utilis.UserNameExist(userLogin.Username) {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	json.NewEncoder(w).Encode(map[string]string{
		// 		"usernameError": "Username does not exist",
		// 	})
		// 	return
		// }

		if err := utilis.CompareHashedPassword(userLogin.Username, &hashedPassword, &userID); err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"usernameError": "Username Does not exist ",
			})
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userLogin.Password)) != nil {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]string{
				"passwordError": "Invalid password",
			})
			return
		}
		sessionID := uuid.New().String()

		expiration := time.Now().Add(24 * time.Hour)
		err := utilis.StoreSession(sessionID, userID, expiration)
		if err != nil {
			http.Error(w, "Failed to store the session", http.StatusInternalServerError)
		}
		cookie := http.Cookie{Name: "session_token", Value: sessionID, Expires: expiration}
		http.SetCookie(w, &cookie)

		// Redirect the user to the posts page
		// http.Redirect(w, r, "/posts", http.StatusSeeOther)
		// row := utils.DB.QueryRow("SELECT password FROM users WHERE email = ?", user.Email)
		// var hashedPassword string
		// if err := row.Scan(&hashedPassword); err != nil {
		// 	http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		// 	return
		// }

		// if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		// 	http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		// 	return
		// }

		// Create a session (for simplicity, use UUID as session ID)
		// sessionID := uuid.New().String()
		// http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionID, Path: "/", HttpOnly: true})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
