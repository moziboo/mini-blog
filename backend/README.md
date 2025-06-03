# Mini-Blog Backend

This is the backend API for the Mini-Blog application, built with Go and Echo framework.

## Project Structure

```
backend/
├── cmd/             # Command-line entry points
│   ├── api/         # API server entry point
│   └── cli/         # CLI tool entry point
├── internal/        # Application code
│   ├── api/         # API handlers and routes
│   ├── db/          # Database operations
│   └── models/      # Data models
└── pkg/             # Reusable packages
    └── logger/      # Logging utilities
```

## Getting Started

### Prerequisites

- Go 1.23 or higher

### Running the API Server

```bash
cd apps/backend
go run cmd/api/main.go
```

The server will start on http://localhost:8080 by default.

### CLI Usage

The CLI tool provides commands for database management:

```bash
# Initialize the database schema
cd apps/backend
go run cmd/cli/main.go migrate

# Seed the database with sample data
go run cmd/cli/main.go seed
```

## API Endpoints

### Posts

- `GET /api/posts` - Get all posts
- `GET /api/posts/:id` - Get a specific post by ID
- `POST /api/posts` - Create a new post
- `PUT /api/posts/:id` - Update a post
- `DELETE /api/posts/:id` - Delete a post

## Development

### Building the Application

```bash
cd apps/backend
go build -o bin/api ./cmd/api
go build -o bin/cli ./cmd/cli
```

### Running Tests

```bash
cd apps/backend
go test ./...
``` 