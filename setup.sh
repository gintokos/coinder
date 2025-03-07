#!/bin/bash

# Auto-deployment script for Docker application
# This script will extract Docker images and deploy the application

# Exit on any error
set -e

echo "========== Docker Deployment Script =========="
echo "Starting deployment process..."

# Check for Docker and Docker Compose
if ! command -v docker &> /dev/null; then
    echo "ERROR: Docker is not installed. Please install Docker first."
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "WARNING: Docker Compose not found. Will attempt to use 'docker compose' instead."
    COMPOSE_CMD="docker compose"
else
    COMPOSE_CMD="docker-compose"
fi

# Extract Docker images from archive
echo "Loading Docker images from archive..."
if [ -f "docker-images.tar" ]; then
    docker load < docker-images.tar
    echo "Docker images loaded successfully!"
else
    echo "ERROR: docker-images.tar not found!"
    exit 1
fi

# Check if docker-compose.yml exists
if [ ! -f "docker-compose.yml" ]; then
    echo "ERROR: docker-compose.yml not found!"
    exit 1
fi

# List loaded images
echo "Loaded Docker images:"
docker images | grep myapp

# Start containers
echo "Starting containers..."
$COMPOSE_CMD up -d

# Check container status
echo "Container status:"
$COMPOSE_CMD ps

echo "Deployment completed successfully!"
echo "=============================================="