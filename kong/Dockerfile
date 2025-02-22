FROM alpine
RUN apk --no-cache add curl
RUN apk --no-cache add gettext
COPY . /
RUN chmod +x /configure_kong.sh
ENTRYPOINT ["/configure_kong.sh"]
