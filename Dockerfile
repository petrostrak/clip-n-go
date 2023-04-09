# Build stage
FROM golang:1.20.3-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main *.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8001
ENTRYPOINT [ "/app/main" ]