.PHONY: start-local build clean build-backend build-frontend help

# Start local development environment
start-local:
	@echo "Starting backend and frontend services..."
	npx concurrently \
		"go run cmd/app/main.go" \
		"cd frontend/v1.0 && npm run dev"

# Makefile for building both backend and frontend

# Force OS to be Windows for testing
OS := windows

# Define OS-specific variables
ifeq ($(OS),windows)
	EXECUTABLE_EXT := .exe
	OS_FLAG := windows
	RM := rm -rf
	MKDIR := mkdir -p
	CP := cp -r
	SEP := /
else
	EXECUTABLE_EXT :=
	OS_FLAG := linux
	RM := rm -rf
	MKDIR := mkdir -p
	CP := cp -r
	SEP := /
endif

# Paths
CMD_DIR := cmd/app
BUILD_DIR := build
APP_BUILD_DIR := $(BUILD_DIR)/app
FRONTEND_SRC_DIR := frontend/v1.0
FRONTEND_DIST_DIR := $(FRONTEND_SRC_DIR)/dist
FRONTEND_BUILD_DIR := $(BUILD_DIR)/frontend

# Main build command
build: build-backend build-frontend

# Build backend (Go application)
build-backend:
	@echo "Building backend for $(OS)..."
	$(MKDIR) $(APP_BUILD_DIR)
ifeq ($(OS),windows)
	go build -o $(APP_BUILD_DIR)/app.exe ./$(CMD_DIR)
else
	go build -o $(APP_BUILD_DIR)/app ./$(CMD_DIR)
endif
	@echo "Backend build completed"

# Build frontend (Vite project)
build-frontend:
	@echo "Building frontend..."
	@cd $(FRONTEND_SRC_DIR) && npm install && npm run build
	$(MKDIR) $(FRONTEND_BUILD_DIR)
	$(CP) $(FRONTEND_DIST_DIR)/* $(FRONTEND_BUILD_DIR)/
	@echo "Frontend build completed"

# Clean build artifacts
clean:
	@echo "Cleaning build directories..."
	$(RM) $(BUILD_DIR)
	@echo "Clean completed"

# Help command
help:
	@echo "Available commands:"
	@echo "  make build         - Build both backend and frontend (default: Windows)"
	@echo "  make build OS=linux - Build for Linux"
	@echo "  make build-backend - Build only the backend Go application"
	@echo "  make build-frontend - Build only the frontend Vite project"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make start-local   - Start local development environment with hot-reloading"