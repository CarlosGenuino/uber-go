# Estágio de construção
FROM golang:1.23-alpine AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia o código fonte para o contêiner
COPY . .

# Baixa as dependências
RUN go mod download

# Compila a aplicação
RUN go build -o uber-go ./cmd/main.go

# Estágio de execução
FROM alpine:latest

# Instala dependências necessárias
RUN apk --no-cache add ca-certificates

# Define o diretório de trabalho
WORKDIR /app

# Copia o binário compilado do estágio de construção
COPY --from=builder /app/uber-go .

# Copia o arquivo .env para o contêiner
COPY .env .

# Expõe a porta 8080
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./uber-go"]