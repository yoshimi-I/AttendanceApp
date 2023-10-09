include .env

DSN=$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)

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
