package utilis

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
// type DB interface {

// }
func OpenDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {

		return nil, err
	}

	return db, nil
}

func ExecuteSQLFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(string(data))
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTables() {
	ExecuteSQLFile("schema.sql")
}

// func InsertUser(email, username, passwordHash string) error {
// 	_, err := DB.Exec(`INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)`, email, username, passwordHash)
// 	return err
// }

// InsertPost inserts a new post into the database.
func InsertPost(userID int, title, content,preview string) error {
	_, err := DB.Exec(`INSERT INTO posts (user_id, title, content, preview) VALUES (?, ?, ?, ?)`, userID, title, content,preview)
	return err
}

// InsertComment inserts a new comment into the database.
func InsertComment(postID, userID int, content string) error {
	_, err := DB.Exec(`INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)`, postID, userID, content)
	return err
}

// InsertCategory inserts a new category into the database.
func InsertCategory(name string) error {
	_, err := DB.Exec(`INSERT INTO categories (name) VALUES (?)`, name)
	return err
}

// AssociatePostWithCategory associates a post with a category.
func AssociatePostWithCategory(postID, categoryID int) error {
	_, err := DB.Exec(`INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)`, postID, categoryID)
	return err
}

// InsertLike inserts a like or dislike into the database.
func InsertLike(userID, postID, commentID int, isLike bool) error {
	_, err := DB.Exec(`INSERT INTO likes (user_id, post_id, comment_id, is_like) VALUES (?, ?, ?, ?)`, userID, postID, commentID, isLike)
	return err
}



// GetPostsByCategory retrieves posts associated with a specific category.
func GetPostsByCategory(categoryName string) (*sql.Rows, error) {
	return DB.Query(`
        SELECT p.* FROM posts p
        JOIN post_categories pc ON p.id = pc.post_id
        JOIN categories c ON pc.category_id = c.id
        WHERE c.name = ?
    `, categoryName)
}

// GetCommentsForPost retrieves all comments for a specific post.
func GetCommentsForPost(postID int) (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM comments WHERE post_id = ?`, postID)
}

// GetPostsLikedByUser retrieves posts liked by a specific user.
func GetPostsLikedByUser(userID int) (*sql.Rows, error) {
	return DB.Query(`
        SELECT p.* FROM posts p
        JOIN likes l ON p.id = l.post_id
        WHERE l.user_id = ? AND l.is_like = 1
    `, userID)
}

// GetPostsCreatedByUser retrieves posts created by a specific user.
func GetPostsCreatedByUser(userID int) (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM posts WHERE user_id = ?`, userID)
}

// GetLikeDislikeCounts retrieves like and dislike counts for a specific post.
func GetLikeDislikeCounts(postID int) (likeCount, dislikeCount int, err error) {
	err = DB.QueryRow(`
        SELECT
            (SELECT COUNT(*) FROM likes WHERE post_id = ? AND is_like = 1) AS like_count,
            (SELECT COUNT(*) FROM likes WHERE post_id = ? AND is_like = 0) AS dislike_count
    `, postID, postID).Scan(&likeCount, &dislikeCount)
	return
}
func InsertUser(email, username, passwordHash string) error {
	_, err := DB.Exec(`INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)`, email, username, passwordHash)
	return err
}
