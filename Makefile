.PHONY: test build clean fmt vet lint install-deps example-basic example-upload

test:
	go test -v -race -coverprofile=coverage.out ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

build:
	go build -v ./...

clean:
	rm -f coverage.out coverage.html
	go clean

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

install-deps:
	go mod download
	go mod tidy

example-basic:
	go run examples/basic/main.go

example-upload:
	go run examples/upload/main.go

help:
	@echo "Available targets:"
	@echo "  test           - Run tests with race detector and coverage"
	@echo "  coverage       - Generate HTML coverage report"
	@echo "  build          - Build the package"
	@echo "  clean          - Clean build artifacts"
	@echo "  fmt            - Format code"
	@echo "  vet            - Run go vet"
	@echo "  lint           - Run golangci-lint"
	@echo "  install-deps   - Download and tidy dependencies"
	@echo "  example-basic  - Run basic example"
	@echo "  example-upload - Run upload example"
