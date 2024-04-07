BINARY_NAME=xkcd
TARGET=cmd/xkcd

all: build

build:
	@echo "Сборка $(BINARY_NAME)..."
	@go build -o ./$(BINARY_NAME) ./$(TARGET)
	@echo "Сборка $(BINARY_NAME) завершена!"

run:
	@go build -o ./$(BINARY_NAME) ./$(TARGET)
	@./$(BINARY_NAME)