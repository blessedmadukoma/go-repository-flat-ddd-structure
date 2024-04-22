FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o main .

FROM alpine:latest  

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/.env .

EXPOSE 8000

CMD ["./main"]
