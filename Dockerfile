FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o kvstore .

FROM scratch

COPY --from=builder /app/kvstore /kvstore


EXPOSE 8080

ENTRYPOINT ["/kvstore"]