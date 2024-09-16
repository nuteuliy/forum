package routes

import (
	"FORUM/handlers"
	"net/http"
)

func SetupRoutes() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/post", handlers.PostDetailsHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/create-post", handlers.CreatePostHandler)
	http.HandleFunc("/view-posts", handlers.ViewPostsHandler)
	// Protected routes
	// http.Handle("/posts", utils.AuthenticateMiddleware(http.HandlerFunc(handlers.CreatePostHandler)))
	// http.Handle("/comments", utils.AuthenticateMiddleware(http.HandlerFunc(handlers.CreateCommentHandler)))
}
