FROM golang:1.22.2 AS builder

WORKDIR /app/DataReceiver

COPY DataReceiver/go.mod DataReceiver/go.sum ./

RUN go mod download

COPY DataReceiver/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

RUN ls -la /app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/DataReceiver/app .

COPY .env .

CMD ["./app"]