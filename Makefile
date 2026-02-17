.PHONY: build run test lint docker-build docker-run clean

# Default binary output.
BIN := server
PORT ?= 8080

## build: Compile the Go binary.
build:
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o bin/$(BIN) ./cmd/server

## run: Run the server locally.
run:
	PORT=$(PORT) go run ./cmd/server

## test: Run all tests with race detection.
test:
	go test -race -count=1 ./...

## lint: Run golangci-lint (must be installed separately).
lint:
	golangci-lint run ./...

## docker-build: Build the Docker image.
docker-build:
	docker build -t go-web-template:local .

## docker-run: Run the Docker container.
docker-run: docker-build
	docker run --rm -p $(PORT):$(PORT) -e PORT=$(PORT) go-web-template:local

## clean: Remove build artifacts.
clean:
	rm -rf bin/
