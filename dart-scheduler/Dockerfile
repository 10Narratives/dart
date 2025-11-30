FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dart-scheduler ./dart-scheduler/cmd/dart-scheduler

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/dart-scheduler .
CMD ["./dart-scheduler", "--config", "/root/config.yaml"]