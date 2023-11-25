# ベースイメージを指定
FROM node:20.9.0-alpine

WORKDIR /app

# 必要なパッケージをインストール
RUN apk update && \
    apk add --no-cache \
    git \
    openssh-client \
    curl \
    bash \
    jq \
    && rm -rf /var/cache/apk/*

# use cache
COPY ./node/package.json ./node/package-lock.json ./
RUN npm ci

COPY ./node .
CMD [ "npm", "run", "dev"]
