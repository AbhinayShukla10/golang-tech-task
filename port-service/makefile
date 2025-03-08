# Define variables
APP_NAME = port-service
BUILD_DIR = bin
DOCKER_IMAGE = port-service
DOCKER_CONTAINER = port-service-container
PORT = 8080

# Default target
.PHONY: all
all: build

# Install dependencies
.PHONY: deps
deps:
	go mod tidy

# Build the Go application
.PHONY: build
build: deps
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/main.go

# Run the application locally
.PHONY: run
run: build
	./$(BUILD_DIR)/$(APP_NAME)

# Run tests
.PHONY: test
test:
	go test ./... -cover

# Lint the code
.PHONY: lint
lint:
	golangci-lint run ./...

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Build Docker image
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container
.PHONY: docker-run
docker-run:
	docker run -p $(PORT):$(PORT) --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

# Stop Docker container
.PHONY: docker-stop
docker-stop:
	docker stop $(DOCKER_CONTAINER) || true
	docker rm $(DOCKER_CONTAINER) || true

# Clean up build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	go clean

# Clean Docker images & containers
.PHONY: docker-clean
docker-clean: docker-stop
	docker rmi $(DOCKER_IMAGE) || true
	docker system prune -f

# Run everything (build, test, run)
.PHONY: all
all: clean build test run
