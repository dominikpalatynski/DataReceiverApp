FROM golang:1.22.2 AS builder

WORKDIR /app/DeviceMenager

COPY DeviceMenager/go.mod DeviceMenager/go.sum ./

RUN go mod download

COPY DeviceMenager/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

RUN ls -la /app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/DeviceMenager/app .

CMD ["./app"]