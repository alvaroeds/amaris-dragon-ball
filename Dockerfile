FROM golang:1.24.3-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod tidy

COPY . .
RUN go build -o main ./cmd

FROM alpine:latest

# Certificados CA necesarios para HTTPS hacia dragonball-api.com
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
