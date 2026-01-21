# Step 1: Build the Go App
FROM golang:alpine AS builder

WORKDIR /app

# Copy dependency files first (Caching layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the executable named 'server'
RUN go build -o server main.go

# Step 2: Run the App (Small Image)
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/server .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./server"]