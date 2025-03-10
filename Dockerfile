# Build stage
FROM golang:1.23.4 AS builder

# Set the working directory to /app in the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy && go mod download

# Copy source code (including cmd folder containing main.go)
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main ./

# Final stage (minimal image for the application)
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
RUN chmod +x /app/main

# Copy the .env file to the container (make sure it's in the same directory as the binary)
COPY --from=builder /app/.env .

# Add CA certificates to allow HTTPS requests
RUN apk --no-cache add ca-certificates


# Set the entrypoint to the binary file
CMD ["./app/main"]