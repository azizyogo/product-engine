FROM golang:1.22.6-alpine AS builder
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files haven't changed
RUN go mod download
# Copy app files
COPY . .
# Build the Go app
RUN go build -o main ./cmd
# Expose port 8080
EXPOSE 8080
# Start the app
CMD ["./main"]
