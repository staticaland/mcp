# Build the MCP server
build:
    go build -o bin/mcp-server .

# Run the MCP server
run:
    go run .

# Clean build artifacts
clean:
    rm -rf bin/

# Download and tidy dependencies
deps:
    go mod download
    go mod tidy

# Format Go code
fmt:
    go fmt ./...

# Run Go vet for static analysis
vet:
    go vet ./...

# Run tests
test:
    go test ./...

# Run tests with coverage
test-coverage:
    go test -cover ./...

# Install the binary locally
install:
    go install .

# Show help
help:
    @just --list