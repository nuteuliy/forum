package handlers

import (
	"FORUM/utilis"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"time"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		template := template.Must(template.ParseFiles("templates/create_post.html"))
		template.Execute(w, nil)
		return
	}
	if r.Method == http.MethodPost {
		var preview string
		cookie, err := r.Cookie("session_token")
		if err != nil {

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		userID, err := utilis.GetSession(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		var PostData struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		if err := json.NewDecoder(r.Body).Decode(&PostData); err != nil {
			http.Error(w, "Invalid input", http.StatusInternalServerError)
			return
		}

		// words := strings.Fields(PostData.Content)

		// preview := strings.Join(words[:30], " ") + "..."
		if len(PostData.Content) < 100 {
			w.Header().Set("Content Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"countError": "Minimum number of words must be 100",
			})
			return
		} else {
			words := strings.Fields(PostData.Content)

			preview = strings.Join(words[:30], " ") + "..."
		}
		// Insert the new post ito the databasen
		err = utilis.InsertPost(userID, PostData.Title, PostData.Content, preview)
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
func PostDetailsHandler(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("id")
	post, err := utilis.GetPostById(postID)
	if err != nil {
		http.Error(w, "POST WAS DELETED", http.StatusInternalServerError)
		return
	}
	cookie, err := r.Cookie("session_token")
	userID, err := utilis.GetSession(cookie.Value)
	isAuthorized := userID != 0
	data := struct {
		Title          string
		FullContent    string
		PreviewContent string
		UserID         int
		CreatedAt      time.Time
		IsAuthorized   bool
	}{
		Title:          post.Title,
		FullContent:    post.Content,
		PreviewContent: post.Preview, // Function to generate preview
		UserID:         post.UserID,
		CreatedAt:      post.CreatedAt,
		IsAuthorized:   isAuthorized,
	}
	tmpl := template.Must(template.ParseFiles("templates/post_details.html"))
	tmpl.Execute(w, data)
}
