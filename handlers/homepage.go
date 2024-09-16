package handlers

import (
	"FORUM/utilis"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := utilis.GetAllPosts()
	if err != nil {

		http.Error(w, "Unable to load posts", http.StatusInternalServerError)
		return 
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, posts)
}
