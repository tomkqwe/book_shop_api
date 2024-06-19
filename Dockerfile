FROM golang:1.22-alpine as builder

RUN go version

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apk update && apk add postgresql bash

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY app-config.yaml /app/app-config.yaml
COPY migrations /app/migrations

EXPOSE 8080

CMD ["./app"]
