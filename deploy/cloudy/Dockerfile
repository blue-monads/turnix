# Builder stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the cloudy binary
RUN GOOS=linux go build -o cloudy ./cmd/cloudy

# Final stage - lightweight Alpine
FROM alpine:latest

# Install ca-certificates and runtime GCC support for Turso shared library
RUN apk --no-cache add ca-certificates libgcc

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/cloudy .

# Set the entrypoint
ENTRYPOINT ["./cloudy"]
