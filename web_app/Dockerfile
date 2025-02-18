# Start from a Go image
FROM golang:1.18-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the container's working directory
COPY . .

# Install any dependencies (if necessary)
RUN go mod tidy

# Build the Go app
RUN go build -o web_app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web_app"]
