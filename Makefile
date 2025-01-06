.PHONY: run test build clean generate tidy

# Default target
.DEFAULT_GOAL := run

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test ./... -v

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Generate module
generate:
	@echo "Enter module name:"
	@read module; \
	go run cmd/generator/generate.go init -m $$module

# Generate module with tests and mocks
generate-with-tests:
	@echo "Enter module name:"
	@read module; \
	go run cmd/generator/generate.go init -m $$module -tests -mocks

# Delete module
delete-module:
	@echo "Enter module name to delete:"
	@read module; \
	go run cmd/generator/generate.go delete -m $$module

# Update dependencies
tidy:
	go mod tidy

# Run linter
lint:
	go vet ./...
	golangci-lint run

# Generate mocks for testing
mocks:
	go generate ./...
