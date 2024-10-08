# 1. Build stage
FROM golang:1.22.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /kitchen

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go binaries orders, kitchen, cookhouse
RUN go build -o kms ./kitchen/.

# 2. Final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /kitchen

# Copy the Go binary from the build stage
COPY --from=builder /kitchen/kms .

# Command to run the Go app
CMD ["./kms"]