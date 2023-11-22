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
│   ├── domain
│   │   ├── model
│   │   │   ├── attendance.go
│   │   │   ├── attendanceType.go
│   │   │   ├── user.go
│   │   │   ├── userStatus.go
│   │   │   └── userStatusType.go
│   │   └── repository
│   │       ├── activityRepository.go
│   │       ├── historyRepository.go
│   │       └── userRepository.go
│   ├── go.mod
│   ├── go.sum
│   ├── infrastructure
│   │   ├── connectionDB.go
│   │   ├── di
│   │   │   ├── wire.go
│   │   │   └── wire_gen.go
│   │   ├── orm
│   │   │   └── gorm_model.go
│   │   └── repository
│   │       ├── activityRepositoryImpl.go
│   │       ├── historyRepositoryImpl.go
│   │       └── userRepositoryImpl.go
│   ├── main.go
│   ├── presentation
│   │   ├── controller
│   │   │   ├── activityController.go
│   │   │   ├── historyController.go
│   │   │   └── userController.go
│   │   ├── parameter
│   │   └── router
│   │       ├── middleware
│   │       │   └── cors.go
│   │       └── router.go
│   ├── tmp
│   │   └── main
│   ├── usecase
│   │   ├── activityUsecase.go
│   │   ├── dto
│   │   │   ├── request
│   │   │   │   ├── activityDTO.go
│   │   │   │   └── user.go
│   │   │   └── response
│   │   │       ├── activityDto.go
│   │   │       ├── historyDto.go
│   │   │       └── user.go
│   │   ├── historyUsecase.go
│   │   └── userUsecase.go
│   └── utility
│       ├── errorUtility.go
│       └── timeUtility.go
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
