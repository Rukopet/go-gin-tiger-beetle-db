.PHONY: build run dev test clean mod-tidy

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run: build
	./bin/server

# Run in development mode
dev:
	go run cmd/server/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Download and tidy dependencies
mod-tidy:
	go mod tidy
	go mod download

# Install tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Lint code
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Check code
check: fmt lint test

# Show help
help:
	@echo "Available commands:"
	@echo "  build      - Build the application"
	@echo "  run        - Build and run the application" 
	@echo "  dev        - Run in development mode"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  mod-tidy   - Download and tidy dependencies"
	@echo "  lint       - Lint code"
	@echo "  fmt        - Format code"
	@echo "  check      - Format, lint and test" 