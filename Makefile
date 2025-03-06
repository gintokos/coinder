.PHONY: start-local

# Start local development environment
start-local:
	@echo "Starting backend and frontend services..."
	npx concurrently \
		"cd && backend go run cmd/app/main.go" \
		"cd frontend/v1.0 && npm run dev"