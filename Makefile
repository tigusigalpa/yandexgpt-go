.PHONY: help test test-coverage lint fmt build clean install examples

help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	go test -v -race ./...

test-coverage: ## Run tests with coverage
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linter
	golangci-lint run

fmt: ## Format code
	go fmt ./...
	goimports -w .

vet: ## Run go vet
	go vet ./...

build: ## Build the package
	go build -v ./...

clean: ## Clean build artifacts
	rm -f coverage.out coverage.html
	go clean -cache -testcache

install: ## Install dependencies
	go mod download
	go mod tidy

examples: ## Run all examples (requires env vars)
	@echo "Running basic example..."
	cd examples/basic && go run main.go
	@echo "\nRunning dialogue example..."
	cd examples/dialogue && go run main.go
	@echo "\nRunning options example..."
	cd examples/options && go run main.go

check: fmt vet lint test ## Run all checks

.DEFAULT_GOAL := help
