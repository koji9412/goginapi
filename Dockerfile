FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /app

# 必要なツールのインストール
RUN go install github.com/air-verse/air@v1.52.3

# モジュールファイルをコピーし、依存関係をインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# airの設定ファイルをコピー
COPY .air.toml .air.toml

# アプリケーションを実行
CMD ["air", "-c", ".air.toml"]
