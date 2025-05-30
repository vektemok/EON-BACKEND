# Название бинарника
BINARY_NAME=myapp

# Путь к main-файлу (если он не в корне)
MAIN_FILE=main.go



.PHONY: run build clean test lint


sgen:
	sqlc generate

	
## Запуск приложения c локальной конфигурацией
run-local:
	# CONFIG_PATH=./config/local.yaml go run cmd/api-server/main.go
	CONFIG_PATH=./config/local.yaml air

	
## Сборка бинарника
build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

## Очистка бинарника
clean:
	rm -f $(BINARY_NAME)

## Запуск тестов
test:
	go test ./...

## Линтинг (нужен golangci-lint)
lint:
	golangci-lint run

## Установка зависимостей (если нужно)
deps:
	go mod tidy
