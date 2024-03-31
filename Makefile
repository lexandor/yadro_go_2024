BINARY_NAME=myapp

all: build

build:
	@echo "Сборка $(BINARY_NAME)..."
	@go build -o ./$(BINARY_NAME) ./stammer/
	@echo "Сборка $(BINARY_NAME) завершена!"

