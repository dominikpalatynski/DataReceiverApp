FROM golang:1.22.2 AS builder

WORKDIR /app/GolangApp

COPY GolangApp/go.mod GolangApp/go.sum ./

RUN go mod download

COPY GolangApp/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app/GolangApp/app ./

CMD ["./app"]