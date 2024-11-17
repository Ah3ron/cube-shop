# Build stage
FROM golang:1.23.3 AS builder

WORKDIR /app/backend

# Copy go mod files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY backend/ .

# Build the application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.20.3

WORKDIR /app

# Install required runtime dependencies (if needed)
RUN apk add --no-cache ca-certificates

# Copy the binary from builder
COPY --from=builder /app/backend/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Expose port
EXPOSE 3000

# Run the application
ENTRYPOINT ["./main"]
