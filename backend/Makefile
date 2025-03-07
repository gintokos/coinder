APP_NAME := app
MAIN_PATH := ./cmd/app/main.go
BUILD_DIR := ./build/app

ifeq ($(OS),LINUX)
	BINARY_NAME := $(APP_NAME)
	GOOS := linux
	GOARCH := amd64
else ifeq ($(OS),WINDOWS)
	BINARY_NAME := $(APP_NAME).exe
	GOOS := windows
	GOARCH := amd64
else ifeq ($(OS),DARWIN)
	BINARY_NAME := $(APP_NAME)
	GOOS := darwin
	GOARCH := amd64
else
	BINARY_NAME := $(APP_NAME)
	GOOS := $(shell go env GOOS)
	GOARCH := $(shell go env GOARCH)
endif

.PHONY: all clean build test

all: clean build

clean:
	@echo "Очистка..."
	@rm -rf $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)

build:
	@echo "Сборка для $(GOOS)/$(GOARCH)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Готово! Бинарный файл: $(BUILD_DIR)/$(BINARY_NAME)"

test:
	@echo "Запуск тестов..."
	@go test -v ./...