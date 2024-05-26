# Use the official Golang image as a base image
FROM golang:1.21.5-alpine

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum to the container
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/go-openai

# Expose the port that the Go application will run on
EXPOSE 8000

# Command to run the Go application
CMD ["./main"]