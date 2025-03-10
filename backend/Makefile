APP_NAME := app
PARSER_NAME := parser
MAIN_PATH := ./cmd/app/main.go
PARSER_PATH := ./cmd/parser/main.go
BUILD_DIR := ./build

ifeq ($(OS),LINUX)
	APP_BINARY_NAME := $(APP_NAME)
	PARSER_BINARY_NAME := $(PARSER_NAME)
	GOOS := linux
	GOARCH := amd64
else ifeq ($(OS),WINDOWS)
	APP_BINARY_NAME := $(APP_NAME).exe
	PARSER_BINARY_NAME := $(PARSER_NAME).exe
	GOOS := windows
	GOARCH := amd64
else ifeq ($(OS),DARWIN)
	APP_BINARY_NAME := $(APP_NAME)
	PARSER_BINARY_NAME := $(PARSER_NAME)
	GOOS := darwin
	GOARCH := amd64
else
	APP_BINARY_NAME := $(APP_NAME)
	PARSER_BINARY_NAME := $(PARSER_NAME)
	GOOS := $(shell go env GOOS)
	GOARCH := $(shell go env GOARCH)
endif

.PHONY: all clean build build-app build-parser test

all: clean build

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)/app
	@mkdir -p $(BUILD_DIR)/parser

build: build-app build-parser

build-app:
	@echo "Building app for $(GOOS)/$(GOARCH)..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/app/$(APP_BINARY_NAME) $(MAIN_PATH)
	@echo "Done! Binary file: $(BUILD_DIR)/app/$(APP_BINARY_NAME)"

build-parser:
	@echo "Building parser for $(GOOS)/$(GOARCH)..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/parser/$(PARSER_BINARY_NAME) $(PARSER_PATH)
	@echo "Done! Binary file: $(BUILD_DIR)/parser/$(PARSER_BINARY_NAME)"

test:
	@echo "Starting tests..."
	@go test -v ./...