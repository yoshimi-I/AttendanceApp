-- +goose Up
-- +goose StatementBegin
ALTER TABLE Attendances ADD year INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Attendances DROP COLUMN year;
-- +goose StatementEnd
