BINARY_NAME=app.exe

all: test lint swagger build

start: build run

build: swagger
	@echo "Building application..."
	@go build -o $(BINARY_NAME) ./cmd/app

run:
	@./$(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@golangci-lint run

swagger:
	@echo "Generating Swagger docs..."
	@swag init -g cmd/app/main.go

.DEFAULT_GOAL := build
