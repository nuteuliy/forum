package handlers

import (
	"FORUM/utilis"
	"html/template"
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		template := template.Must(template.ParseFiles("templates/create_post.html"))
		template.Execute(w, nil)
		return
	}
	if r.Method == http.MethodPost {
		
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		userID, err := utilis.GetSession(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		title := r.FormValue("title")
		content := r.FormValue("content")

		// Insert the new post into the database
		err = utilis.InsertPost(userID, title, content)
		if err != nil {
			http.Error(w, "Error creating post", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func ViewPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := utilis.GetAllPosts()
	if err != nil {
		http.Error(w, "Unable to load posts", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/view_posts.html"))
	tmpl.Execute(w, posts)
}
