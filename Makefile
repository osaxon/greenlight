# Variables
APP_NAME := greenlight
BUILD_DIR := ./bin
CMD_DIR := ./cmd/api
ENTRY_POINT := $(CMD_DIR)

.PHONY: build
build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(ENTRY_POINT)

.PHONY: run
run:
	@echo "Running the application..."
	@go run $(CMD_DIR)

.PHONY: dev
dev:
	@echo "Starting the application with Air..."
	@air

.PHONY: clean
clean:
	@echo "Cleaning up the build directory..."
	@rm -rf $(BUILD_DIR)

migrate-up:
	@go migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up