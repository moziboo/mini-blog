package models

import (
	"time"
)

// Post represents a blog post
type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPost creates a new post with the given title and content
func NewPost(title, content string) *Post {
	now := time.Now()
	return &Post{
		Title:     title,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}
} 