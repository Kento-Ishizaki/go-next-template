FROM golang:1.21-alpine as dev

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

COPY go/go.mod go/go.sum ./
RUN go mod download

# ホットリロード用ツール
RUN go install github.com/cosmtrek/air@latest
# マイグレーション用ツール
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["air", "-c", ".air.toml"]

FROM golang:1.21-alpine as builder

WORKDIR /app

RUN apk update && \
    apk add --no-cache \
    git \
    && rm -rf /var/cache/apk/*

COPY go/go.mod go/go.sum ./
RUN go mod download

COPY go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/binary

FROM scratch as prod
WORKDIR /app
COPY --from=builder /app/binary .
CMD ["./binary"]
