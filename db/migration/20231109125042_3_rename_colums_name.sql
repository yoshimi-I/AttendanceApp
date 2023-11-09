-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_status_types
    CHANGE COLUMN action_type user_status_type VARCHAR(100) NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_status_types
    CHANGE COLUMN user_status_type action_type VARCHAR(100) NOT NULL;

-- +goose StatementEnd
