
# Stage 1: Build the Go application
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Download dependencies first for better build caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source
COPY . .

# Build a standalone Linux executable
RUN CGO_ENABLED=0 GOOS=linux go build -o /api .

# Stage 2: Run the application
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /api /app/api

EXPOSE 8080

CMD ["/app/api"]

