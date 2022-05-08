FROM golang:1.18 AS builder

WORKDIR /app

# Copy and downloading the modules first and seperately allows docker to cache this layer and speed up the build when no changes were made
COPY go.mod go.sum ./
RUN go mod download


COPY . .

# Compile production ready binary with no runtime dependencies and stripped debug information
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server/main.go


FROM alpine:latest

# For using external services(like firebase) with https
RUN apk add --no-cache git ca-certificates

COPY --from=builder /app .

# Use non root user with least privileges
RUN addgroup -S sandbox && adduser -S -G sandbox sandbox && chown -R sandbox:sandbox .
USER sandbox

CMD ["./server"]
