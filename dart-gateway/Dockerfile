# dart-gateway/Dockerfile
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dart-gateway ./dart-gateway/cmd/dart-gateway

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/dart-gateway .
EXPOSE 8080
CMD ["./dart-gateway", "--config", "/root/config.yaml"]