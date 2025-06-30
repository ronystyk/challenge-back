# Etapa 1: Build
FROM golang:1.24 AS builder

WORKDIR /app

# Copiar go.mod y go.sum y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario para Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Etapa 2: Imagen final
FROM alpine:latest

WORKDIR /root/

# Copiar el binario desde el builder
COPY --from=builder /app/main .

# Exponer el puerto en el que corre tu app
EXPOSE 8080

# Comando por defecto
CMD ["./main"]