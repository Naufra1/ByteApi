FROM golang:1.22.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /byteApi

COPY ./db/main.db /app/db/main.db

FROM debian:bookworm-slim

COPY --from=builder /byteApi /byteApi
COPY --from=builder /app/db/main.db /app/db/main.db

WORKDIR /app

EXPOSE 8080

CMD ["/byteApi"]
