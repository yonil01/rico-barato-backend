# Use a slim Alpine Linux base image
FROM golang:alpine AS builder

# Set working directory for the build stage
WORKDIR /app

# Copy your Go source code (replace "your-app" with your actual directory)
COPY your-app/. .

# Install dependencies using Go modules (adjust if using vendor directory)
RUN go mod download

# Build the Go binary (adjust the output name if needed)
RUN go build -o main

# Use a smaller, multi-stage image for the final container
FROM alpine:3.16

# Copy the built binary from the builder stage
COPY --from=builder /app/main /app/main

# Set the working directory for the final container
WORKDIR /app

# Expose the port your Fiber application listens on (adjust as needed)
EXPOSE 8080

# Command to run your Fiber application
CMD ["/app/main"]
