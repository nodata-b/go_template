# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
WORKDIR /go/src/github.com/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/app
# パッケージのインポート
RUN go get -u golang.org/x/tools/cmd/goimports
