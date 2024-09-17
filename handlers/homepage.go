package handlers

import (
	"FORUM/utilis"
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := utilis.GetAllPosts()
	fmt.Println(posts)

	if err != nil {

		http.Error(w, "Unable to load posts", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, posts)
}
