FROM golang:alpine AS build
COPY . .
ENV GOPATH=""
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -a -ldflags="-w -s" -o api

FROM scratch
COPY --from=build /go/api /api
ENTRYPOINT ["/api"]
