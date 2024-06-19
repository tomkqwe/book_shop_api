FROM golang:1.22-alpine as builder

RUN go version

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# install psql
RUN apk update && \
apk add postgresql

# install bash
RUN apk add --no-cache --upgrade bash

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh
COPY wait-for-postgres.sh /wait-for-postgres.sh

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go
EXPOSE 8080

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
