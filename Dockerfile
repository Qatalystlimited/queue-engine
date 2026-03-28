FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o queue-engine .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/queue-engine .
EXPOSE 50051
CMD ["./queue-engine"]
