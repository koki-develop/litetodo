# See: https://github.com/mattn/go-sqlite3/blob/master/_example/simple/Dockerfile

FROM golang:1.21-alpine3.19 AS builder
ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -o app

# ---

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/app ./

CMD ["./app"]
