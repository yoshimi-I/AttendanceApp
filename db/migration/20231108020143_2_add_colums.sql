-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN firebase_id VARCHAR(255) UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN firebase_id;

-- +goose StatementEnd
