-- +goose Up

-- +goose StatementBegin
INSERT INTO User (user_id, name, email) VALUES
                                            (1, '山田太郎', 'taro.yamada@example.com'),
                                            (2, '鈴木花子', 'hanako.suzuki@example.com');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO Attendances (attendance_id, user_id, attendance_date, clock_in_time, clock_out_time, notes) VALUES
                                                                                                            (1, 1, '2023-10-08', '2023-10-08 09:00:00', '2023-10-08 18:00:00', '出勤日'),
                                                                                                            (2, 2, '2023-10-08', '2023-10-08 10:00:00', '2023-10-08 19:00:00', '遅刻');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO Breaks (break_id, attendance_id, start_time, end_time) VALUES
                                                                       (1, 1, '2023-10-08 12:00:00', '2023-10-08 13:00:00'),
                                                                       (2, 2, '2023-10-08 12:30:00', '2023-10-08 13:30:00');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO PrayerTimes (prayer_id, attendance_id, fajr, dhuhr, asr, maghrib, isha) VALUES
                                                                                        (1, 1, '2023-10-08 05:00:00', '2023-10-08 12:00:00', '2023-10-08 15:00:00', '2023-10-08 17:00:00', '2023-10-08 19:00:00'),
                                                                                        (2, 2, '2023-10-08 05:30:00', '2023-10-08 12:30:00', '2023-10-08 15:30:00', '2023-10-08 17:30:00', '2023-10-08 19:30:00');
-- +goose StatementEnd

-- +goose Down
-- ここにテーブルからデータを削除するSQLを書いてください。例えば:
-- +goose StatementBegin
DELETE FROM PrayerTimes WHERE prayer_id IN (1, 2);
-- +goose StatementEnd

-- +goose StatementBegin
DELETE FROM Breaks WHERE break_id IN (1, 2);
-- +goose StatementEnd

-- +goose StatementBegin
DELETE FROM Attendances WHERE attendance_id IN (1, 2);
-- +goose StatementEnd

-- +goose StatementBegin
DELETE FROM User WHERE user_id IN (1, 2);
-- +goose StatementEnd
