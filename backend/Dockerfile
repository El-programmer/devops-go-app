# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Install git (needed for go mod sometimes)
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# Now copy the rest of the code
COPY . .

# Build
RUN go build -o main .

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]
