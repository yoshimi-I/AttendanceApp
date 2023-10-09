-- +goose Up

-- +goose StatementBegin
INSERT INTO User (name, email) VALUES
('Yoshimi', 'yoshimi@example.com'),
('Taro', 'taro@example.com');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO Activities (user_id, activity_date, notes) VALUES
(1, '2023-10-10', 'Yoshimi study notes'),
(2, '2023-10-10', 'Taro study notes');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO Play (activity_id, start_time, end_time) VALUES
(1, '2023-10-10 09:00:00', '2023-10-10 10:00:00'),
(2, '2023-10-10 10:00:00', '2023-10-10 11:00:00');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO StudyStartEnd (activity_id, start_time, end_time) VALUES
(1, '2023-10-10 11:00:00', '2023-10-10 12:00:00'),
(2, '2023-10-10 12:00:00', '2023-10-10 13:00:00');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO BreakStartEnd (activity_id, start_time, end_time) VALUES
(1, '2023-10-10 13:00:00', '2023-10-10 14:00:00'),
(2, '2023-10-10 14:00:00', '2023-10-10 15:00:00');
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DELETE FROM BreakStartEnd WHERE activity_id IN (1, 2);
-- +goose StatementEnd

-- +goose StatementBegin
DELETE FROM StudyStartEnd WHERE activity_id IN (1, 2);
-- +goose StatementEnd

-- +goose StatementBegin
DELETE FROM Play WHERE activity_id IN (1, 2);
-- +goose StatementEnd

-- +goose StatementBegin
DELETE FROM Activities WHERE user_id IN (1, 2);
-- +goose StatementEnd

-- +goose StatementBegin
DELETE FROM User WHERE user_id IN (1, 2);
-- +goose StatementEnd
