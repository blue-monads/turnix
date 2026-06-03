# Builder stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the cloudy binary
RUN CGO_ENABLED=0 GOOS=linux go build -o cloudy ./cmd/cloudy

# Final stage - lightweight Alpine
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/cloudy .

# Set the entrypoint
ENTRYPOINT ["./cloudy"]
