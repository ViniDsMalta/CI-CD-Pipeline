# build
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/api

# Final
FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]