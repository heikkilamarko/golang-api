FROM alpine

RUN apk --no-cache add curl

COPY . /

RUN chmod +x /configure_kong.sh

ENTRYPOINT [ "/configure_kong.sh" ]
