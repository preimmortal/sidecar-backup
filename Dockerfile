# pull official base image
FROM alpine:3.15
RUN apk add --no-cache --virtual .run-deps bash rsync sqlite curl  && rm -rf /var/cache/apk/*
COPY entrypoint.sh /
COPY build/
ENTRYPOINT [ "/entrypoint.sh"]