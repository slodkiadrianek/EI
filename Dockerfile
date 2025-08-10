FROM golang:1.24.7 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env ./

RUN GOOS=linux GOARCH=arm GOARM=7 go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 3031

CMD ["./main"]
