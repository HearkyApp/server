FROM golang:1.18 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./server"]