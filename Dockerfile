FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o kvstore .


FROM alpine:latest

RUN adduser -D -u 1000 appuser

RUN mkdir -p /data && chown -R appuser:appuser /data

COPY --from=builder /app/kvstore /kvstore

USER appuser

EXPOSE 8080

ENTRYPOINT ["/kvstore"]
