BINARY=bin/gobank

build:
	@echo "==> Compilando o projeto..."
	@go build -o $(BINARY) ./cmd/main.go

run: build
	@echo "==> Iniciando a API..."
	@./$(BINARY)

test:
	@echo "==> Rodando testes..."
	@go test ./... -v

