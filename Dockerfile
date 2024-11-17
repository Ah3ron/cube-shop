# Build stage
FROM golang:1.23.3 AS builder

WORKDIR /backend

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.20.3

WORKDIR /app

# Install required runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the binary from builder
COPY --from=builder /app/main .

# Set environment variables
ENV DATABASE_URL=""
ENV JWT_SECRET=""

# Expose port
EXPOSE 3000

# Run the application
CMD ["./main"]