FROM golang:1.21-alpine

WORKDIR /app

# ホットリロード用ツールをインストール
RUN go install github.com/cosmtrek/air@latest

# 必要なパッケージをインストール
RUN apk update && \
    apk add --no-cache \
    git \
    openssh-client \
    curl \
    bash \
    jq \
    && rm -rf /var/cache/apk/*

COPY go/go.mod go/go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
