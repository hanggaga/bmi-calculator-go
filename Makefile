# Variables
SERVICE_DIR = ./service
COVERAGE_FILE = coverage.out
BINARY_NAME = bmi_calculator

# Default target
.PHONY: all
all: test-cover run

# Run tests with coverage
.PHONY: test-cover
test-cover:
	@echo "Running tests with coverage..."
	go test -coverprofile=$(COVERAGE_FILE) $(SERVICE_DIR)
	go tool cover -func=$(COVERAGE_FILE)

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	go build -o $(BINARY_NAME) ./app/main.go

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	./$(BINARY_NAME)
