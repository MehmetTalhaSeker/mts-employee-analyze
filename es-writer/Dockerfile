# Builder
FROM --platform=linux/amd64 golang:1.21.4-alpine3.17 as builder

RUN apk add alpine-sdk

ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOARCH="amd64"
ENV GOOS=linux

WORKDIR /build

COPY go.* ./

RUN apk add musl-dev

RUN go mod download

COPY . ./

RUN go build -tags musl -o es-writer .

# Application container
FROM alpine:3.17

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /lib/ld-musl-x86_64.so.1 /lib/

COPY --from=builder /build/es-writer /app/es-writer

CMD ["/app/es-writer"]