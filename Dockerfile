# =========================
# Build stage
# =========================
FROM golang:1.25.5-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server


# =========================
# Runtime stage
# =========================
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/server .

# Optional: copy .env (ONLY for local/dev)
# In production, use environment variables instead
COPY .env .env

# Expose app port
EXPOSE 8080

# Run the server
CMD ["./server"]
