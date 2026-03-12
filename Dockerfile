# Dockerfile

# Stage 1: Build the Go application
FROM golang:1.18 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o server cmd/server/main.go

# Stage 2: Create a lightweight image for running the application
FROM alpine:latest

# Set the current working directory inside the new image
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/server .

# Expose the port that the app runs on
EXPOSE $PORT

# Command to run the executable
CMD ["./server"]
