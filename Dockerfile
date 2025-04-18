# Build stage
FROM golang:1.22-rc-alpine AS builder

WORKDIR /app

# Install required build tools
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary and environment file from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Set environment variables
ENV TZ=Asia/Jakarta

# # Create a non-root user and switch to it
# RUN adduser -D appuser
# RUN chown -R appuser:appuser /app
# USER appuser

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
