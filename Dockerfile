FROM golang:alpine AS builder

ARG CGO_ENABLED=0 
ARG GOOS=linux
RUN go install github.com/google/wire/cmd/wire@latest

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal

RUN wire ./cmd/api
RUN go build -o app ./cmd/api

FROM alpine
LABEL maintainer="Kirill Korhunov <korhunov.kv@edu.spbstu.ru>"
LABEL org.opencontainers.image.source=https://github.com/kir0108/PayShareBackend

WORKDIR /app

COPY --from=builder /build/app .
EXPOSE 8080
CMD ["./app"]
