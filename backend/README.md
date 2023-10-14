## バックエンド

### 使用言語
  - Go言語
### ライブラリ
  - Gorm
    - ORM
  - Goose
    - マイグレーション
  - go-chi
    - https
  - Air
    - ホットリロード
  - swaggo
    - ドキュメント自動生成
### ディレクトリ構成
```
.
├── main.go  # アプリケーションのエントリーポイント
├── config          # 設定ファイルや環境変数の管理
│   └── config.go
├── domain
│   ├── model
│   │   ├── user.go
│   │   ├── attendance.go
│   │   └── ...     # その他のモデル
│   └── repository
│       ├── user_repository.go
│       ├── attendance_repository.go
│       └── ...     # その他のリポジトリインターフェース
├── infrastructure
│   ├── config      # DBや外部サービスの接続設定
│   │   └── connection.go
│   ├── orm
│   │   ├── gorm_model.go    # GORMに関連する設定やマッピング
│   │   └── ...     # 他のORMの設定やマッピング
│   ├── repository
│   │   ├── mysql   # MySQLに関連するリポジトリの実装
│   │   ├── redis   # Redisに関連するリポジトリの実装
│   │   └── ...     # その他のデータソースのリポジトリ実装
│   └── middleware  # ミドルウェアの実装
├── usecase
│   ├── user_usecase.go
│   ├── attendance_usecase.go
│   └── ...         # その他のユースケース
├── router          # ルーティングの設定
│   └── router.go
├── di              # 依存性注入の設定
│   └── di.go
├── go.mod
├── go.sum
└── go.work

```

## 
```bash
go mod init yoshimi-I/AttendanceApp
```

## Dockerで迷子になったら
- まずはディレクトリ構成を把握
```
docker build --target build -t temp-image .
```
- そのあとlsやpwdを押してルートを確認
```
docker run --rm temp-image ls
```
