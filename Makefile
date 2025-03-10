.PHONY: build tag archive package clean all start-local

# Start local development environment
start-local:
	@echo "Starting backend and frontend services..."
	npx concurrently \
		"cd backend && go run cmd/app/main.go" \
		"cd frontend/v1.0 && npm run dev"

# Image names and versions
FRONTEND_IMAGE := myapp-frontend
BACKEND_IMAGE := myapp-backend
NGINX_IMAGE := myapp-nginx
VERSION := 1.0

# Docker Registry tag if needed
REGISTRY :=

# Archive names
DOCKER_ARCHIVE := docker-images.tar
DEPLOY_ARCHIVE := deploy.tar.gz

all: build tag archive package

# Build Docker images
build:
	@echo "Building Docker images..."
	docker build -t $(FRONTEND_IMAGE):$(VERSION) ./frontend/v1.0
	docker build -t $(BACKEND_IMAGE):$(VERSION) ./backend
	docker build -t $(NGINX_IMAGE):$(VERSION) .
	@echo "Docker images built successfully"

# Add tags to images
tag:
	@echo "Tagging Docker images..."
	@if [ -n "$(REGISTRY)" ]; then \
		docker tag $(FRONTEND_IMAGE):$(VERSION) $(REGISTRY)/$(FRONTEND_IMAGE):$(VERSION); \
		docker tag $(BACKEND_IMAGE):$(VERSION) $(REGISTRY)/$(BACKEND_IMAGE):$(VERSION); \
		docker tag $(NGINX_IMAGE):$(VERSION) $(REGISTRY)/$(NGINX_IMAGE):$(VERSION); \
	fi
	@echo "Docker images tagged successfully"

# Save images to archive
archive:
	@echo "Saving Docker images to archive..."
	docker save $(FRONTEND_IMAGE):$(VERSION) $(BACKEND_IMAGE):$(VERSION) $(NGINX_IMAGE):$(VERSION) > $(DOCKER_ARCHIVE)
	@echo "Docker images saved to $(DOCKER_ARCHIVE)"

# Create deployment package archive
package:
	@echo "Creating deployment package..."
	# Create a temporary directory
	mkdir -p ./tmp_deploy
	# Copy files to the temporary directory
	cp $(DOCKER_ARCHIVE) ./tmp_deploy/
	cp docker-compose.yml ./tmp_deploy/
	cp setup.sh ./tmp_deploy/
	cp cleanup.sh ./tmp_deploy/
	cp .env ./tmp_deploy/
	chmod +x ./tmp_deploy/setup.sh
	chmod +x ./tmp_deploy/cleanup.sh
	# Create the deployment archive
	tar -czf $(DEPLOY_ARCHIVE) -C ./tmp_deploy .
	# Remove the temporary directory
	rm -rf ./tmp_deploy
	@echo "Deployment package created: $(DEPLOY_ARCHIVE)"

# Load images from archive (for testing locally)
load:
	@echo "Loading Docker images from archive..."
	docker load < $(DOCKER_ARCHIVE)
	@echo "Docker images loaded successfully"

# Clean local resources
clean:
	@echo "Cleaning up..."
	rm -f $(DOCKER_ARCHIVE) $(DEPLOY_ARCHIVE)
	rm -rf ./tmp_deploy
	docker rmi $(FRONTEND_IMAGE):$(VERSION) $(BACKEND_IMAGE):$(VERSION) $(NGINX_IMAGE):$(VERSION) || true
	@echo "Cleanup completed"

# Help information
help:
	@echo "Available targets:"
	@echo "  build   - Build Docker images"
	@echo "  tag     - Tag Docker images with registry prefix (if REGISTRY is set)"
	@echo "  archive - Save Docker images to archive file $(DOCKER_ARCHIVE)"
	@echo "  package - Create deployment package with images, compose file and scripts"
	@echo "  load    - Load Docker images from archive file (for testing)"
	@echo "  clean   - Remove Docker images and archives"
	@echo "  all     - Run build, tag, archive and package (default target)"
	@echo "  start-local - Start backend and frontend for local development"