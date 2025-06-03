package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// DB represents a database connection
type DB struct {
	*sql.DB
}

// New creates a new database connection
func New(path string) (*DB, error) {
	// Create the directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Check if the database file exists, if not create it
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("failed to create database file: %w", err)
		}
		file.Close()
	}

	// Open the database connection
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{DB: db}, nil
}

// Migrate runs the database migrations
func (db *DB) Migrate() error {
	// Create posts table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create posts table: %w", err)
	}

	return nil
}

// GetPosts retrieves all posts from the database
func (db *DB) GetPosts() (*sql.Rows, error) {
	return db.Query(`
		SELECT id, title, content, created_at, updated_at 
		FROM posts 
		ORDER BY created_at DESC
	`)
}

// GetPost retrieves a post by ID
func (db *DB) GetPost(id int64) (*sql.Row, error) {
	return db.QueryRow(`
		SELECT id, title, content, created_at, updated_at 
		FROM posts 
		WHERE id = ?
	`, id), nil
}

// CreatePost adds a new post to the database
func (db *DB) CreatePost(title, content string) (int64, error) {
	result, err := db.Exec(`
		INSERT INTO posts (title, content) 
		VALUES (?, ?)
	`, title, content)
	if err != nil {
		return 0, fmt.Errorf("failed to create post: %w", err)
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	
	return id, nil
}

// PostExists checks if a post with the given ID exists
func (db *DB) PostExists(id int64) (bool, error) {
	var exists int
	err := db.QueryRow("SELECT 1 FROM posts WHERE id = ?", id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if post exists: %w", err)
	}
	return true, nil
}

// UpdatePost updates a post in the database
func (db *DB) UpdatePost(id int64, title, content string) error {
	_, err := db.Exec(`
		UPDATE posts 
		SET title = ?, content = ?, updated_at = CURRENT_TIMESTAMP 
		WHERE id = ?
	`, title, content, id)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}
	return nil
}

// DeletePost deletes a post from the database
func (db *DB) DeletePost(id int64) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}
	return nil
}

// PostExistsByTitle checks if a post with the given title exists
func (db *DB) PostExistsByTitle(title string) (bool, error) {
	var exists int
	err := db.QueryRow("SELECT 1 FROM posts WHERE title = ?", title).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if post exists by title: %w", err)
	}
	return true, nil
}

// UpdatePostByTitle updates a post in the database by its title
func (db *DB) UpdatePostByTitle(title, content string) error {
	_, err := db.Exec(`
		UPDATE posts 
		SET content = ?, updated_at = CURRENT_TIMESTAMP 
		WHERE title = ?
	`, content, title)
	if err != nil {
		return fmt.Errorf("failed to update post by title: %w", err)
	}
	return nil
}