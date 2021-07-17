# ベースとなるイメージ
FROM golang:1.13.7-buster

# プロジェクトルートに移動する
WORKDIR /go/src/gin-train

# Air インストール
# RUN curl -fLo /go/bin/air https://git.io/linux_air \
#   && chmod +x /go/bin/air
RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

# コンテナ実行時のデフォルトを設定する
# ライブリロードを実行する
CMD air