FROM golang:1.22.1 AS builder
WORKDIR /app
COPY . .
RUN go build -o order-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-service .
CMD ["./order-service"]
