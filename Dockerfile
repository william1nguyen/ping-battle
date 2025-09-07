FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bin/ping-battle ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/ping-battle /app/ping-battle

EXPOSE 8080

CMD ["./ping-battle"]