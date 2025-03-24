# Build Stage
FROM golang:1.24 AS builder

# Create the working directory inside the container
WORKDIR /app

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Move to the directory where `main.go` is located
WORKDIR /app/cmd/server

# Build the Go application
RUN go build -o server .

# Final Stage - Use a minimal base image
FROM debian:latest

# Create the working directory in the runtime container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/cmd/server/server .

# Copy the .env file
COPY .env /app/.env

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./server"]