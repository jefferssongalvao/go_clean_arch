# Stage 1: build
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Copia os arquivos Go
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build do binário principal
RUN go build -o go-clean-api ./cmd/api/main.go
RUN go build -o go-clean-api-migrate ./cmd/migrate/main.go

# Stage 2: final
FROM alpine:3.18

WORKDIR /app

# Copia os binários construídos
COPY --from=builder /app/go-clean-api .
COPY --from=builder /app/go-clean-api-migrate .

# Copia .env (se necessário)
COPY .env .

# Executável padrão
CMD ["./go-clean-api"]
