# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .


# Development stage
FROM golang:1.21-alpine AS development

WORKDIR /app

# Install air for hot-reloading
RUN go install github.com/cosmtrek/air@latest

# Copy source code
COPY --from=builder /app /app

# Set up the entrypoint for development
ENTRYPOINT ["air"]


# Production stage
FROM alpine:latest AS production

WORKDIR /app

# Install required packages
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy environment files (if any)
COPY --from=builder /app/.env* ./

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
