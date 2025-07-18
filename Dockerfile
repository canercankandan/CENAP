# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Modülleri önceden kopyala ve indir
COPY go.mod go.sum ./
RUN go mod download

# Tüm kaynak kodunu kopyala
COPY . .

# Uygulamayı derle
RUN go build -o main ./cmd/web

# Run stage
FROM alpine:latest

WORKDIR /app

# Gerekli dosyaları kopyala
COPY --from=builder /app/main .
COPY static ./static
COPY templates ./templates
COPY data.json .
COPY orders.json .

# Render için port belirtme - Render kendi PORT environment variable'ını kullanır
# EXPOSE 8080
# EXPOSE 8081

# Uygulamayı başlat
CMD ["./main"] 