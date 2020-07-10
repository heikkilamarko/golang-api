# build stage
FROM golang:alpine AS build-env

WORKDIR /app

COPY ./src/ ./

RUN go build -ldflags="-s -w" -o products-api

# runtime stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/products-api .
ENTRYPOINT ["./products-api"]
