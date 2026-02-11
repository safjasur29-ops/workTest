FROM golang:1.23.4 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o workTest

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/workTest .

EXPOSE 8080

CMD ["./workTest"]