# pull official base image
FROM alpine:3.15

RUN apk add --no-cache --virtual .run-deps bash rsync sqlite curl  && rm -rf /var/cache/apk/*
RUN mkdir /app/

ENV PUID=0
ENV PGID=0
ENV ARCH=$TARGETARCH

COPY entrypoint.sh /
COPY sidecar-backup.${ARCH} /app/sidecar-backup

ENTRYPOINT [ "/entrypoint.sh"]