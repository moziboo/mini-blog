package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/stephen.garrett/mini-blog/backend/internal/db"
	"github.com/stephen.garrett/mini-blog/backend/pkg/logger"
)

func main() {
	// Initialize logger
	logger := logger.New()

	// Define command line flags
	dbPath := flag.String("db", "./mini-blog.db", "Database file path")
	
	// Define subcommands
	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	
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
		
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: cli [options] <command>")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nCommands:")
	fmt.Println("  migrate    Run database migrations")
} 