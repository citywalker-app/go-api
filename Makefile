include .env.dev
include .env.prod

lint:
	@golangci-lint run

run-prod:
	@echo "Building..."
	@go build -o main cmd/api/main.go
	@echo "Running..."
	@ENV_FILE=.env.prod ./main &

# Run the application
run-dev:
	@ENV_FILE=.env.dev go run cmd/api/main.go

# Test the application
test:
	@echo "Clearing test cache..."
	@go clean -testcache
	@echo "Testing..."
	@ENV_FILE=../.env.dev go test ./... -coverprofile=coverage.out
