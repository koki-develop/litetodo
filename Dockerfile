FROM golang:1.21-alpine3.19 AS builder
WORKDIR /app
ENV LITESTREAM_VERSION=v0.3.13
ENV CGO_ENABLED=1

# Install litestream
ADD https://github.com/benbjohnson/litestream/releases/download/$LITESTREAM_VERSION/litestream-$LITESTREAM_VERSION-linux-amd64.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

# Install dependencies
RUN apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY main.go ./
RUN go build -o app

# ---

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/app ./
COPY --from=builder /usr/local/bin/litestream /usr/local/bin/litestream
COPY run.sh ./

CMD ["/app/run.sh"]
