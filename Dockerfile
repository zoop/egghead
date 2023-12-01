# Use an official Go runtime as a parent image
FROM golang:latest AS builder

# Set the working directory in the container
WORKDIR /go/src/app

# Copy only the go.mod and go.sum files to take advantage of layer caching
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o index .

# Use a lightweight base image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /go/src/app/index .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
# CMD ["ls", "-la"]
CMD ["./index"]
