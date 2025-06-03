package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stephen.garrett/mini-blog/backend/internal/api"
	"github.com/stephen.garrett/mini-blog/backend/internal/db"
	"github.com/stephen.garrett/mini-blog/backend/pkg/logger"
)

func main() {
	// Parse command line flags
	addr := flag.String("addr", ":1323", "HTTP server address")
	dbPath := flag.String("db", "./db/mini-blog.db", "Database file path")
	flag.Parse()

	// Initialize logger
	logger := logger.New()
	logger.Info("Starting API server")

	// Initialize database
	database, err := db.New(*dbPath)
	if err != nil {
		logger.Fatal("Failed to connect to database", "error", err)
	}
	defer database.Close()

	// Initialize Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Setup API routes
	api.RegisterRoutes(e, database, logger)

	// Start server
	go func() {
		if err := e.Start(*addr); err != nil {
			logger.Info("Shutting down server", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", "error", err)
	}
	
	logger.Info("Server has been shutdown")
} 