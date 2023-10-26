-- +goose Up
-- +goose StatementBegin
RENAME TABLE Users TO users;
-- +goose StatementEnd
-- +goose StatementBegin
RENAME TABLE Attendance_types TO attendance_types;
-- +goose StatementEnd
-- +goose StatementBegin
RENAME TABLE Attendances TO attendances;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
RENAME TABLE users TO Users;
-- +goose StatementEnd
-- +goose StatementBegin
RENAME TABLE attendance_types TO Attendance_types;
-- +goose StatementEnd
-- +goose StatementBegin
RENAME TABLE attendances TO Attendances;
-- +goose StatementEnd
