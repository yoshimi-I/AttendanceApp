# ビルドステージ
FROM oven/bun:1.0.0 as build

WORKDIR /frontend

# package.json と package-lock.json をコピー
COPY package.json package-lock.json ./

# 依存関係のインストール
RUN bun install

# 他のファイルとディレクトリをコピー
COPY . .

# bunを使用してNext.jsのビルド
RUN bun next build

# 実行ステージ
FROM oven/bun:1.0.0 as final

WORKDIR /frontend


# ポートの公開
EXPOSE 3000

# 起動コマンドの設定
CMD ["bun", "start"]
