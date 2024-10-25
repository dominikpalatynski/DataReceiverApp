# Użyj obrazu Go jako podstawowego
FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/ ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app/app .

CMD ["./app"]