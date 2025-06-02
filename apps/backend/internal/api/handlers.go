package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stephen.garrett/mini-blog/backend/internal/db"
	"github.com/stephen.garrett/mini-blog/backend/internal/models"
	"github.com/stephen.garrett/mini-blog/backend/pkg/logger"
)

// Handler represents the API handlers
type Handler struct {
	DB     *db.DB
	Logger logger.Logger
}

// NewHandler creates a new API handler
func NewHandler(db *db.DB, logger logger.Logger) *Handler {
	return &Handler{
		DB:     db,
		Logger: logger,
	}
}

// RegisterRoutes registers the API routes
func RegisterRoutes(e *echo.Echo, db *db.DB, logger logger.Logger) {
	h := NewHandler(db, logger)

	// Group API routes under /api
	api := e.Group("/api")

	// Posts routes
	api.GET("/posts", h.GetPosts)
	api.GET("/posts/:id", h.GetPost)
	api.POST("/posts", h.CreatePost)
	api.PUT("/posts/:id", h.UpdatePost)
	api.DELETE("/posts/:id", h.DeletePost)
}

// GetPosts returns all posts
func (h *Handler) GetPosts(c echo.Context) error {
	rows, err := h.DB.GetPosts()
	if err != nil {
		h.Logger.Error("Failed to get posts", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get posts")
	}
	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var p models.Post
		var createdAt, updatedAt string
		
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &createdAt, &updatedAt); err != nil {
			h.Logger.Error("Failed to scan post row", "error", err)
			continue
		}
		
		p.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		p.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)
		posts = append(posts, p)
	}

	return c.JSON(http.StatusOK, posts)
}

// GetPost returns a post by ID
func (h *Handler) GetPost(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	row, err := h.DB.GetPost(id)
	if err != nil {
		h.Logger.Error("Failed to get post", "error", err, "id", id)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get post")
	}

	var p models.Post
	var createdAt, updatedAt string
	
	if err := row.Scan(&p.ID, &p.Title, &p.Content, &createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}
		h.Logger.Error("Failed to scan post row", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get post")
	}
	
	p.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	p.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return c.JSON(http.StatusOK, p)
}

// CreatePost creates a new post
func (h *Handler) CreatePost(c echo.Context) error {
	type request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	req := new(request)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if req.Title == "" || req.Content == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title and content are required")
	}

	id, err := h.DB.CreatePost(req.Title, req.Content)
	if err != nil {
		h.Logger.Error("Failed to create post", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create post")
	}

	post := models.NewPost(req.Title, req.Content)
	post.ID = id

	return c.JSON(http.StatusCreated, post)
}

// UpdatePost updates an existing post
func (h *Handler) UpdatePost(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	type request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	req := new(request)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if req.Title == "" || req.Content == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title and content are required")
	}

	exists, err := h.DB.PostExists(id)
	if err != nil {
		h.Logger.Error("Failed to check if post exists", "error", err, "id", id)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update post")
	}

	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	if err := h.DB.UpdatePost(id, req.Title, req.Content); err != nil {
		h.Logger.Error("Failed to update post", "error", err, "id", id)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update post")
	}

	// Get the updated post
	row, err := h.DB.GetPost(id)
	if err != nil {
		h.Logger.Error("Failed to get updated post", "error", err, "id", id)
		return echo.NewHTTPError(http.StatusInternalServerError, "Post updated but failed to retrieve it")
	}

	var p models.Post
	var createdAt, updatedAt string
	
	if err := row.Scan(&p.ID, &p.Title, &p.Content, &createdAt, &updatedAt); err != nil {
		h.Logger.Error("Failed to scan updated post row", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Post updated but failed to retrieve it")
	}
	
	p.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	p.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return c.JSON(http.StatusOK, p)
}

// DeletePost deletes a post
func (h *Handler) DeletePost(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	exists, err := h.DB.PostExists(id)
	if err != nil {
		h.Logger.Error("Failed to check if post exists", "error", err, "id", id)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete post")
	}

	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	if err := h.DB.DeletePost(id); err != nil {
		h.Logger.Error("Failed to delete post", "error", err, "id", id)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete post")
	}

	return c.NoContent(http.StatusNoContent)
} 