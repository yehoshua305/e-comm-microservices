# Build Stage
FROM golang:1.23-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git

# Set environment variables for ARM64 build
ENV GOARCH=arm64
ENV GOOS=linux
RUN go build -o main main.go

# Run Stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]