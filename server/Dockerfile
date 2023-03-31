# ベースとなるDockerイメージ指定
FROM golang:1.20.2

# コンテナ内に作業ディレクトリを作成
RUN mkdir -p /go/src/github.com/aocm/vue-go-spa-sample
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/github.com/aocm/vue-go-spa-sample
# ホストのファイルをコンテナの作業ディレクトリに移行
# ADD ./ /go/src/github.com/aocm/vue-go-spa-sample

# モジュールインポート
# RUN cd /go/src/github.com/aocm/vue-go-spa-sample
# RUN go get -d

# godocインストール
# RUN go get golang.org/x/tools/cmd/godoc