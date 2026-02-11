FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o workTest

FROM debian:bookworm-20240110-slim

WORKDIR /app

COPY --from=builder /app/workTest .

EXPOSE 8080

CMD ["./workTest"]