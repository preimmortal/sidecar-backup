FROM --platform=$BUILDPLATFORM golang:1.18.4-alpine3.16 AS build
WORKDIR /src
COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -v -o . ./...

FROM alpine:3.16

RUN apk add --no-cache --virtual .run-deps bash rsync sqlite curl  && rm -rf /var/cache/apk/*
RUN mkdir /app/

ENV PUID=0
ENV PGID=0

COPY entrypoint.sh /entrypoint.sh

COPY --from=build sidecar-backup /usr/bin/sidecar-backup

ENTRYPOINT [ "/entrypoint.sh"]