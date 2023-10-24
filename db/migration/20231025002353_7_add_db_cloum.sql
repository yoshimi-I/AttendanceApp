-- +goose Up
-- +goose StatementBegin
ALTER TABLE Attendances ADD end_time TIMESTAMP;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE Attendances CHANGE actual_time start_time TIMESTAMP;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE Attendances ADD date VARCHAR(15);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Attendances DROP COLUMN end_time;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE Attendances CHANGE start_time actual_time TIMESTAMP;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE Attendances DROP COLUMN date;
-- +goose StatementEnd
