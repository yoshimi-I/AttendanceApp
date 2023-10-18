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
├── README.md
├── api
│   ├── config
│   │   └── config.go
│   ├── controller
│   │   ├── activityController.go ルーティングでここのコントローラを呼ぶ
│   │   └── historyController.go
│   ├── di
│   │   ├── wire.go
│   │   └── wire_gen.go
│   ├── domain
│   │   ├── model // ビジネスモデル
│   │   │   └── model.go 
│   │   └── repository // DB操作の処理をinterfaceで定義
│   │       ├── activityRepository.go 
│   │       └── historyRepository.go
│   ├── go.mod
│   ├── go.sum
│   ├── infrastructure 
│   │   ├── connection.go
│   │   ├── orm
│   │   │   └── gorm_model.go // DBアクセス時に使用
│   │   └── repository // domain層のrepositoryを実装する
│   │       ├── activityRepositoryImpl.go
│   │       └── historyRepositoryImpl.go
│   ├── main.go
│   ├── router
│   │   └── router.go  //ここでルーティング
│   ├── tmp
│   │   └── main
│   └── usecase //今回のアプリケーションのビジネスロジックはここに書く
│       ├── activityUsecase.go
│       ├── dto // model層のデータをdtoを通して書き換える
│       │   ├── activityResponseDto.go
│       │   └── historyResponseDto.go
│       └── historyUsecase.go
└── dockerfile

```
```mermaid
sequenceDiagram

    participant Client
    participant Router
    participant Controller
    participant Usecase
    participant Repository
    participant Database

    Client->>Router: リクエスト (例: GET /activity)

    Router->>Controller: 適切なControllerへリクエストを転送

    Controller->>Usecase: 関連するビジネスロジックを呼び出し

    Usecase->>Repository: データをリクエスト

    Repository->>Database: データを取得/保存

    Database-->>Repository: データを返す
    note left of Repository: 返り値: domain層のmodel

    Repository-->>Usecase: データを返す
    note left of Usecase: 返り値: ResponseDataDTO

    Usecase-->>Controller: 処理されたデータを返す

    Controller-->>Router: Clientへの応答
    note left of Router: 返り値: json

    Router-->>Client: 応答 (例: アクティビティデータ)



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
