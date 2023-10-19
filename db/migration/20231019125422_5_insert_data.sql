-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- Users sample data
INSERT INTO Users (name, email) VALUES
                                    ('John Doe', 'john@example.com'),
                                    ('Jane Smith', 'jane@example.com');



-- Attendances sample data
INSERT INTO Attendances (user_id, attendance_type, notes, actual_time) VALUES
                                                                           (1, 1, 'Johns attendance note', '2023-10-20 09:00:00'),
(1, 2, 'Johns break note', '2023-10-20 12:00:00'),
                                                                           (2, 1, 'Janes attendance note', '2023-10-20 09:15:00');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DELETE FROM Attendances WHERE id IN (1,2,3);
DELETE FROM Users WHERE id IN (1,2);
DELETE FROM Users WHERE id IN (1,2);
