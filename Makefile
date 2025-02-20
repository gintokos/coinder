.PHONY: start-local

start-local:
	@echo "Starting backend and frontend services..."
	npx concurrently \
		"go run cmd/app/main.go" \
		"cd frontend/v1.0 && npm run dev"