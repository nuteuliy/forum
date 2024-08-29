package routes

import (
	"FORUM/handlers"
	"net/http"
	"text/template"
)

func SetupRoutes() {
	fs := http.FileServer(http.Dir("static"))
	
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	// Protected routes
	// http.Handle("/posts", utils.AuthenticateMiddleware(http.HandlerFunc(handlers.CreatePostHandler)))
	// http.Handle("/comments", utils.AuthenticateMiddleware(http.HandlerFunc(handlers.CreateCommentHandler)))
}
