FROM golang:1.18-alpine

# 変数を環境変数にセット
ENV APP_ENV=dev

# install packages
RUN apk update && \
    apk add --update --no-cache \
    tzdata \
    bash \
    make \
    # delveのinstallを
    gcc \
    musl-dev \
    # 日本時間に修正
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo "Asia/Tokyo" > /etc/timezone \
    && apk del tzdata


WORKDIR /bc_app

#  開発時に必要なパッケージをインストール (普遍的なもの)
RUN go mod init github.com/toshi0228/blockchain \
    # Airをインストール
    && go install github.com/cosmtrek/air@v1.40.4 \
    && go install github.com/go-delve/delve/cmd/dlv@latest


# ファイルをホストからコンテナにコピー
COPY . .

# 依存パッケージの更新
RUN go mod tidy


# airコマンドでGoファイルを起動
WORKDIR /bc_app/src/cmd
CMD ["air", "-c", "../../.air.toml"]

