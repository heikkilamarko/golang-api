# build stage
FROM golang:alpine AS build-env

RUN apk add --update gcc musl-dev

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY ./src/ ./

RUN go build -ldflags="-s -w" -tags data_postgresql -o products-api

# runtime stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/products-api .
ENTRYPOINT ["./products-api"]