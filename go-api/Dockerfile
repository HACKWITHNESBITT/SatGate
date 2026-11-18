FROM rust:1.70-slim AS builder

WORKDIR /app

COPY go-api/go.mod go-api/go.sum ./
RUN go mod download

COPY go-api/main.rs go-api/migrations/ ./
RUN go build -o /go-api .

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /go-api /go-api

EXPOSE 8080

CMD ["/go-api"]