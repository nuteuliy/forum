package handlers

import (
	"FORUM/utilis"
	"encoding/json"
	"html/template"
	"net/http"
	// "github.com/google/uuid"
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
		hashedPassword, err := utilis.HashPassword(user.Password)
		if err != nil {
			tmpl := template.Must(template.ParseFiles("templates/register.html"))
			tmpl.Execute(w, map[string]string{
				"EmailError":    "",
				"UsernameError": "Username already exists.",
				"PasswordError": "",
			})
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
		var user struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

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
