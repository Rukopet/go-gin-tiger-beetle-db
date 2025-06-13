.PHONY: build run dev test clean mod-tidy docker-build docker-up docker-down docker-logs docker-clean docker-dev tigerbeetle-format tigerbeetle-logs

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

# Build Docker image
docker-build:
	docker-compose build

# Start services
docker-up:
	docker-compose up -d

# Stop services
docker-down:
	docker-compose down

# Show logs
docker-logs:
	docker-compose logs -f

# Restart services
docker-restart:
	docker-compose restart

# Clean all Docker resources
docker-clean:
	docker-compose down --volumes --rmi all

# Build and start services
docker-dev:
	docker-compose up --build

# Format TigerBeetle data file
tigerbeetle-format:
	docker run --rm -v tigerbeetle-data:/data ghcr.io/tigerbeetle/tigerbeetle:latest format --cluster=0 --replica=0 --replica-count=1 /data/0_0.tigerbeetle

# Show TigerBeetle logs
tigerbeetle-logs:
	docker-compose logs -f tigerbeetle

# Show help
help:
	@echo "Available commands:"
	@echo "  build               - Build the application"
	@echo "  run                 - Build and run the application" 
	@echo "  dev                 - Run in development mode"
	@echo "  test                - Run tests"
	@echo "  clean               - Clean build artifacts"
	@echo "  mod-tidy            - Download and tidy dependencies"
	@echo "  lint                - Lint code"
	@echo "  fmt                 - Format code"
	@echo "  check               - Format, lint and test"
	@echo "  docker-build        - Build Docker image"
	@echo "  docker-up           - Start services"
	@echo "  docker-down         - Stop services"
	@echo "  docker-logs         - Show logs"
	@echo "  docker-restart      - Restart services"
	@echo "  docker-clean        - Clean all Docker resources"
	@echo "  docker-dev          - Build and start services"
	@echo "  tigerbeetle-format  - Format TigerBeetle data file"
	@echo "  tigerbeetle-logs    - Show TigerBeetle logs" 