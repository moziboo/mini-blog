package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/stephen.garrett/mini-blog/backend/internal/db"
	"github.com/stephen.garrett/mini-blog/backend/pkg/logger"
)

func main() {
	// Initialize logger
	logger := logger.New()

	// Define command line flags
	dbPath := flag.String("db", "./db/mini-blog.db", "Database file path")
	
	// Define subcommands
	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	
	// Import command flags
	dataDir := importCmd.String("dir", "./data", "Directory containing markdown files")
	
	// Parse the main command
	flag.Parse()
	
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Connect to database
	database, err := db.New(*dbPath)
	if err != nil {
		logger.Fatal("Failed to connect to database", "error", err)
	}
	defer database.Close()

	// Handle subcommands
	switch os.Args[1] {
	case "migrate":
		migrateCmd.Parse(os.Args[2:])
		logger.Info("Running migrations")
		if err := database.Migrate(); err != nil {
			logger.Fatal("Migration failed", "error", err)
		}
		logger.Info("Migrations completed successfully")
		
	case "import":
		importCmd.Parse(os.Args[2:])
		logger.Info("Importing markdown files", "directory", *dataDir)
		if err := importMarkdownFiles(*dataDir, database, logger); err != nil {
			logger.Fatal("Import failed", "error", err)
		}
		logger.Info("Import completed successfully")
		
	default:
		printUsage()
		os.Exit(1)
	}
}

// importMarkdownFiles imports all markdown files from the specified directory
func importMarkdownFiles(dirPath string, database *db.DB, logger logger.Logger) error {
	// Ensure the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dirPath)
	}

	// Walk through the directory
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only process markdown files
		if !strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
			return nil
		}

		logger.Info("Processing file", "file", path)

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		// Extract title from filename (remove extension)
		title := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		
		// Check if a post with this title already exists
		exists, err := database.PostExistsByTitle(title)
		if err != nil {
			return fmt.Errorf("failed to check if post exists: %w", err)
		}

		if exists {
			logger.Info("Post already exists, updating", "title", title)
			if err := database.UpdatePostByTitle(title, string(content)); err != nil {
				return fmt.Errorf("failed to update post: %w", err)
			}
		} else {
			logger.Info("Creating new post", "title", title)
			if _, err := database.CreatePost(title, string(content)); err != nil {
				return fmt.Errorf("failed to create post: %w", err)
			}
		}

		return nil
	})
}

func printUsage() {
	fmt.Println("Usage: cli [options] <command>")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nCommands:")
	fmt.Println("  migrate    Run database migrations")
	fmt.Println("  import     Import markdown files from a directory")
} 