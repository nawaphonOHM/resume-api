SHELL := /bin/bash
export PATH := $(PATH):/snap/bin:$(HOME)/go/bin
.DEFAULT_GOAL := help

.PHONY: all
all: test-unit vet build

.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build binary and verify compilation of all packages
	go build -v ./...

.PHONY: run
run: ## Run the application
	go run ./cmd/resume-api

.PHONY: test
test: ## Run the full unit and integration test suite with race detection
	go test -race -v -tags="testcontainers,realmongo" ./...

.PHONY: test-unit
test-unit: ## Run unit tests with race detection
	go test -race -v ./...

.PHONY: test-integration-tc
test-integration-tc: ## Run Testcontainers-backed integration tests with race detection
	go test -race -v -tags=testcontainers ./...

.PHONY: test-integration-real
test-integration-real: ## Run real MongoDB integration tests with race detection (requires MONGODB_URI)
	go test -race -v -tags=realmongo ./...

.PHONY: test-integration
test-integration: ## Run combined integration tests (Testcontainers & real Mongo) with race detection
	go test -race -v -tags="testcontainers,realmongo" ./...

.PHONY: test-coverage
test-coverage: ## Run tests with race detection and HTML coverage report
	go test -race -tags="testcontainers,realmongo" -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: vet
vet: ## Run go vet analysis
	go vet -tags="testcontainers,realmongo" ./...

.PHONY: lint
lint: ## Run golangci-lint (or go vet if golangci-lint not installed)
	@which golangci-lint > /dev/null 2>&1 && golangci-lint run --build-tags="testcontainers,realmongo" || (echo "golangci-lint not found in PATH; running go vet ./..." && go vet -tags="testcontainers,realmongo" ./...)

.PHONY: tidy
tidy: ## Tidy and verify Go module dependencies
	go mod tidy
	go mod verify

.PHONY: clean
clean: ## Clean build artifacts and test coverage files
	rm -rf bin tmp coverage.out coverage.html profile.out
	@echo "Cleaned build and test artifacts"
