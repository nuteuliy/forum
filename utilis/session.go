package utilis

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int
	UserID    int
	Title     string
	Content   string
	CreatedAt time.Time
	// UpdatedAt time.Time
}

func StoreSession(sessionID string, userID int, expires time.Time) error {
	// Use the global DB connection
	_, err := DB.Exec(`
        INSERT INTO sessions (session_id, user_id, expires_at)
        VALUES (?, ?, ?)`, sessionID, userID, expires)
	return err
}
func GetSession(sessionID string) (int, error) {
	var userID int
	err := DB.QueryRow(`
        SELECT user_id FROM sessions
        WHERE session_id = ? AND expires_at > CURRENT_TIMESTAMP`, sessionID).Scan(&userID)
	if err == sql.ErrNoRows {
		return 0, nil // No session found
	}
	if err != nil {
		return 0, err // Error occurred
	}
	return userID, nil
}
func GetAllPosts() ([]Post, error) {
	rows, err := DB.Query(`SELECT id, user_id, title, content, created_at 
		FROM posts 
		ORDER BY created_at DESC ;`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.UserID,  &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
