.PHONY: setup start build clean frontend backend dev-backend dev-frontend import-data migrate dev

# Setup both frontend and backend
setup: setup-frontend setup-backend

# Start both services in development mode
dev:
	@echo "Starting backend server in background..."
	@cd backend && go run cmd/api/main.go & echo $$! > .backend.pid
	@echo "Starting frontend dev server..."
	@cd frontend && npm run dev
	@echo "Stopping backend server..."
	@kill `cat .backend.pid` && rm .backend.pid || true

# Setup frontend
setup-frontend:
	cd frontend && npm install

# Setup backend
setup-backend:
	cd backend && go mod tidy

# Run backend migration
migrate:
	cd backend && go run cmd/cli/main.go migrate

# Import sample data
import-data:
	cd backend && go run cmd/cli/main.go import -dir="../_data"

# Run backend in development mode
dev-backend:
	cd backend && go run cmd/api/main.go

# Run frontend in development mode
dev-frontend:
	cd frontend && npm run dev

# Build both frontend and backend
build: build-frontend build-backend

# Build frontend for production
build-frontend:
	cd frontend && npm run build

# Build backend binaries
build-backend:
	mkdir -p bin
	cd backend && go build -o ../bin/api cmd/api/main.go
	cd backend && go build -o ../bin/cli cmd/cli/main.go

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf frontend/dist/
	rm -f .backend.pid