# Stage 1: Build
FROM golang:1.25.7-bookworm AS builder

WORKDIR /app/backend

# Install build dependencies
RUN apt update && apt install -y git curl build-essential

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build statically for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# Stage 2: Minimal runtime
FROM golang:1.25.7-bookworm

WORKDIR /app/backend

# Copy the binary from the builder
COPY --from=builder /app/backend/main .

# Make sure it's executable
RUN chmod +x main

EXPOSE 8080

# Use absolute path for safety
CMD ["/app/backend/main"]
