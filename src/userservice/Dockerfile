# Build Stage
FROM golang:1.23-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go build -o main main.go

# Run Stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]