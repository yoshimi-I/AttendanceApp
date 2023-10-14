include .env

DSN=$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)

# Dockerを立ち上げる
# 全て立ち上げる
docker-all d:
	docker-compose up -d
# フォアグラウンドで立ち上げる
docker-all fg:
	docker-compose up
# dbとapiだけ立ち上げる
docker-db-api:
	docker-compose up db api
# Dockerを停止する
docker-stop:
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
