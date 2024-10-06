# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go binary
RUN go build -o server main.go

# Stage 2: Run the Go application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .
COPY data/ ./data/

# Expose the port the app runs on
EXPOSE 8080

# Run the Go server
CMD ["./server"]
