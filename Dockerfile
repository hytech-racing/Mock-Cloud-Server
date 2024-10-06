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

# ENV Variables
ARG AWS_REGION
ARG AWS_S3_RUN_BUCKET
ARG AWS_ACCESS_KEY
ARG AWS_SECRET_KEY
ENV AWS_REGION=${AWS_REGION}
ENV AWS_S3_RUN_BUCKET=${AWS_S3_RUN_BUCKET}
ENV AWS_ACCESS_KEY=${AWS_ACCESS_KEY}
ENV AWS_SECRET_KEY=${AWS_SECRET_KEY}

# Expose the port the app runs on
EXPOSE 8080

# Run the Go server
CMD ["./server"]
