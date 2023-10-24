include .env

DSN=$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)

# Dockerを立ち上げる
# 全て立ち上げる
all-d:
	docker-compose up -d
# フォアグラウンドで立ち上げる
all-fg:
	docker-compose up
# dbとapiだけ立ち上げる
db-api:
	docker-compose up db backend
# Dockerを停止する
stop:
	docker-compose stop

# マイグレーションの新しいバージョンを作成する
# 例) make migrate-create name=1_add_tables_and_sample_data
migrate-create:
	cd backend/api && goose -dir=../../db/migration create $(name) sql

# マイグレーションを適用する
migrate-up:
	cd backend/api && goose -dir=../../db/migration mysql "$(DSN)" up

# マイグレーションを1つ戻す
migrate-down:
	cd backend/api && goose -dir=../../db/migration mysql "$(DSN)" down

# 現在のマイグレーションの状態を表示する
migrate-status:
	cd backend/api && goose -dir=../../db/migration mysql "$(DSN)" status

# DIを行う
di:
	cd backend/api/di && wire gen