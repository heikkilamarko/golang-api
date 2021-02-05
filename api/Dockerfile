# build stage
FROM golang:1.16rc1-alpine3.13 AS build-env

WORKDIR /app

COPY ./ ./

RUN go build -ldflags="-s -w" -o products-api

# runtime stage
FROM alpine:3.13

COPY --from=build-env /app/products-api /usr/bin

ENTRYPOINT ["/usr/bin/products-api"]
