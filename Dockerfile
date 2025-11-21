# builder
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git build-base

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goravel .

# final
FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/goravel /app/goravel

COPY .env.prod /app/.env

EXPOSE 8080

CMD ["./goravel"]
