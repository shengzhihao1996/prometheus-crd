FROM alpine

COPY bin/app /app/app

COPY config /etc/app/config

RUN echo -e "https://mirrors.aliyun.com/alpine/v3.8/main\nhttps://mirrors.aliyun.com/alpine/v3.8/community" > /etc/apk/repositories \
    && apk add --no-cache curl \
    && chmod +x /app/app

WORKDIR /app

USER root

CMD ["./app"]
