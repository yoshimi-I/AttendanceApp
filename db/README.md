## データベース周り
### 使用技術
  - データベース(RDS)
    - mysql
  - マイグレーションライブラリ
    - Goose(Go言語のライブラリ)
### ER図
  - ./samples/er.md参照
### 手順
  1. Dockerを立ち上げるタイミングでテーブルを生成します
  2. データの挿入はマイグレーションライブラリを使って行います

### マイグレーションの流し方
  1. ./Makefileを参照
  2. Makefileのmigrate-createを流す
  3. ./db/migrationにマイグレーションのsqlファイルができるのでそこに書きたいDDLを記載
  4. Makefileに戻りmigrate-upを実行
  5. 戻したい場合はmigrate-downを実行


  