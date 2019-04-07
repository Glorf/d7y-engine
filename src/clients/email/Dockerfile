FROM golang:1.12.2 as builder
COPY / /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:latest
MAINTAINER Michał Bień
COPY --from=builder /build/email /go/email
ENTRYPOINT /go/email -sender-server $SENDER_SERVER -sender-port $SENDER_PORT
