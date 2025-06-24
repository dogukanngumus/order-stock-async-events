# Stage 1: Build
FROM golang:1.23-alpine AS build

WORKDIR /app

# Modül dosyalarını kopyala ve bağımlılıkları indir
COPY go.mod go.sum ./
RUN go mod download

# Kaynak kodları kopyala
COPY . .

# Statik olarak build et (CGO_ENABLED=0 ile)
RUN CGO_ENABLED=0 GOOS=linux go build -o stock-service .

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

# HTTPS/TLS için gerekli sertifikalar
RUN apk --no-cache add ca-certificates

# Build edilmiş binary'yi kopyala
COPY --from=build /app/stock-service .

# Servisin dışarıya açılacağı port
EXPOSE 8080

# Entry point
ENTRYPOINT ["./stock-service"]
