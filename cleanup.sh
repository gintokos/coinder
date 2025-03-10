#!/bin/bash

# Cleanup script for Docker application
# This script will stop all containers and clean up resources but preserve volumes

# Exit on any error
set -e

echo "========== Docker Cleanup Script =========="
echo "Starting cleanup process..."

# Check for Docker and Docker Compose
if ! command -v docker &> /dev/null; then
    echo "ERROR: Docker is not installed. Cannot perform cleanup."
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "WARNING: Docker Compose not found. Will attempt to use 'docker compose' instead."
    COMPOSE_CMD="docker compose"
else
    COMPOSE_CMD="docker-compose"
fi

# Check if docker-compose.yml exists
if [ -f "docker-compose.yml" ]; then
    echo "Stopping containers (preserving volumes)..."
    $COMPOSE_CMD down  # Removed -v flag to preserve volumes
    echo "Containers stopped and removed successfully. Volumes preserved."
else
    echo "WARNING: docker-compose.yml not found! Will try to find and remove containers manually."
    
    # Find and remove containers by image name pattern
    CONTAINERS=$(docker ps -a --filter "name=myapp" --format "{{.ID}}")
    if [ -n "$CONTAINERS" ]; then
        echo "Found containers to remove. Stopping and removing..."
        docker stop $CONTAINERS 2>/dev/null || true
        docker rm $CONTAINERS 2>/dev/null || true
        echo "Containers removed."
    else
        echo "No matching containers found."
    fi
fi

# Remove Docker images and containers
echo "Removing Docker containers and images..."

# Stop and remove all containers first to prevent conflicts
echo "Stopping all running containers..."
docker ps -aq | xargs -r docker stop
docker ps -aq | xargs -r docker rm

# Force remove application images 
echo "Removing application images..."
docker images | grep "myapp-" | awk '{print $3}' | xargs -r docker rmi -f || true

# Remove postgres image if exists
echo "Removing postgres image..."
docker rmi postgres:14-alpine 2>/dev/null || true

# Aggressive cleanup of all dangling images by ID
echo "Removing dangling images..."
docker images -f "dangling=true" -q | xargs -r docker rmi -f 2>/dev/null || true

# Prune system but exclude volumes
echo "Pruning unused Docker resources (excluding volumes)..."
docker system prune -f --volumes=false  # Added --volumes=false to preserve volumes

echo "Docker cleanup completed. Volumes have been preserved."

# Remove archives and files
echo "Removing deployment files..."
rm -f docker-images.tar setup.sh
echo "Files removed."

echo "Do you want to remove all files in the current directory? (y/n)"
read -r response
if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
    echo "Removing all files in the current directory..."
    rm -rf * .[^.]*
    echo "All files removed."
else
    echo "Skipping removal of all files."
fi

echo "Cleanup completed successfully!"
echo "=============================================="