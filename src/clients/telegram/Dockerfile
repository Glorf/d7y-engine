FROM golang:1.12.2-alpine3.9 AS build-env

COPY . /telegram
WORKDIR /telegram

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git && \
    go get ./...  && \
    apk del git && \ 
    go build

FROM alpine
WORKDIR /app
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates 2>/dev/null || true
COPY --from=build-env /telegram/ /app/
