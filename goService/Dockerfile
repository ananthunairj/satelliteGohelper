# ---------- Build Stage ----------
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your source
COPY . .

# Build a static binary (disable CGO)
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o goservice

# ---------- Final Stage ----------
FROM alpine:latest

# Install tiny curl tool for healthcheck (wget is larger)
RUN apk --no-cache add curl

# Set working directory
WORKDIR /app

# Copy only the built binary from builder
COPY --from=builder /app/goservice .

# Expose port
EXPOSE 8080

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl --fail http://localhost:8080/health || exit 1

# Run the binary
CMD ["./goservice"]
