.PHONY: run build test clean

# Variáveis
APP_NAME := my-favorite-movies
CMD_PATH := ./app/cmd/api

# Rodar a aplicação
run:
	go run $(CMD_PATH)/main.go

# Compilar a aplicação
build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

# Rodar testes
test:
	go test ./...

# Limpar binários compilados
clean:
	rm -rf bin/$(APP_NAME)
