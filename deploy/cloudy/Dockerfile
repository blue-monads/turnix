# Builder stage
FROM golang:1.25-bullseye AS builder

WORKDIR /app

# Install build dependencies for CGO (if needed) and certs
RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
		ca-certificates gcc libc6-dev build-essential && \
	rm -rf /var/lib/apt/lists/*

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the cloudy binary
RUN GOOS=linux go build -o cloudy ./cmd/cloudy

# Final stage - lightweight Alpine
FROM ubuntu:24.04

# Install ca-certificates and runtime GCC support for Turso shared library
RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
		ca-certificates libgcc-s1 && \
	rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/cloudy .

# Set the entrypoint
ENTRYPOINT ["./cloudy"]
