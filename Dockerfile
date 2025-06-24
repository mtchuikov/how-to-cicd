FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk update --no-cache && \
    apk add --no-cache tzdata

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o myapp *.go

FROM alpine:3.22.0

RUN apk update --no-cache && \
    apk add --no-cache ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

ENV TZ=UTC
RUN ln -sf /usr/share/zoneinfo/UTC /etc/localtime

WORKDIR /app

COPY --from=builder /build/myapp /app/myapp

EXPOSE 8080

ENTRYPOINT ["./myapp"]
