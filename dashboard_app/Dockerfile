# Use the official Golang image as a build stage
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o dashboard_app .

# Use a minimal base image for the final container
FROM alpine:latest

# Set working directory in the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/dashboard_app .

# Expose the port the app runs on
EXPOSE 8090

# Run the dashboard application
CMD ["./dashboard_app"]
