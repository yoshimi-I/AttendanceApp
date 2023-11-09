-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN user_key VARCHAR(255) UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN user_key;
-- +goose StatementEnd
