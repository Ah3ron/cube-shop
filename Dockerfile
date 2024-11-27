# Build stage for backend
FROM golang:1.23.3 AS backend-builder

# Set working directory for backend
WORKDIR /app/backend

# Copy go mod files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy backend source code
COPY backend/ .

# Build the application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Build stage for frontend
FROM node:23 AS frontend-builder

# Set working directory for frontend
WORKDIR /app/frontend

# Copy package files and install dependencies using pnpm
COPY frontend/package.json ./
RUN npm install -g pnpm
RUN pnpm install

# Copy frontend source code
COPY frontend/ .

# Build the frontend application
RUN pnpm run build

# Final stage using Alpine Linux
FROM alpine:3.20.3

# Set working directory
WORKDIR /app

# Install SSL certificates for HTTPS support
RUN apk add --no-cache ca-certificates

# Copy built frontend assets
COPY --from=frontend-builder /app/frontend/build ./build
# Copy the backend binary from builder
COPY --from=backend-builder /app/backend/main .

# Make the binary executable
RUN chmod +x ./main

# Expose application port
EXPOSE 8080

# Run the backend application
ENTRYPOINT ["./main"]