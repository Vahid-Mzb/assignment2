FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o kvstore .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/kvstore .
RUN chmod +x ./kvstore

EXPOSE 8080

ENV DATA_PATH=/data/store.json

ENTRYPOINT ["./kvstore"]