FROM golang:1.20-alpine AS build
ENV HOME /opt/recipes

WORKDIR $HOME
COPY . $HOME/

RUN set -x && \
    apk update && \
    apk add --no-cache \
    make

RUN set -x && \
    make OUTPUT_BUILD=/var/lib/app build

FROM alpine:latest AS target
RUN apk add --no-cache tzdata
ENV TZ=Europe/Moscow
COPY --from=build /var/lib/app /usr/local/bin/app
CMD /usr/local/bin/app
