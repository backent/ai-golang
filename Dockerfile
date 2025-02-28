# Use a small base image
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go modules first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application as a static binary
RUN go build -o app .

# Use a minimal base image for the final container
FROM alpine:latest

# Set working directory in the final image
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose the application port
EXPOSE 8022

# Run the application
CMD ["./app"]
