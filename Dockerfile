# Gunakan image base Golang versi terbaru
FROM golang:1.24-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum untuk mengunduh dependencies
COPY go.mod go.sum ./

# Unduh semua dependencies
RUN go mod download

# Copy seluruh kode sumber ke working directory
COPY . .

# Build aplikasi Go
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Gunakan image alpine yang lebih ringan untuk production
FROM alpine:latest

# Install ca-certificates untuk HTTPS
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy binary dari builder
COPY --from=builder /app/main .

# Copy file .env agar bisa dibaca oleh aplikasi
COPY .env .env

# Expose port yang digunakan oleh aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
