FROM golang:1.19 AS builder

WORKDIR /app
ADD . /app

RUN go build -o similigo-api main.go

FROM ubuntu:latest AS launcher
COPY --from=builder /app .
CMD ["./similigo-api", "api"]