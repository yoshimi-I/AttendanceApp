-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_status_types (
                                   user_status_type_id INT PRIMARY KEY,
                                   action_type VARCHAR(100) NOT NULL,
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_statuses (
                               id INT AUTO_INCREMENT PRIMARY KEY,
                               user_id INT,
                               status_id INT,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               FOREIGN KEY (user_id) REFERENCES users(id),
                               FOREIGN KEY (status_id) REFERENCES user_status_types(user_status_type_id)
);
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO user_status_types (user_status_type_id, action_type) VALUES
                                                                    (1, '作業中'),
                                                                    (2, '休憩中'),
                                                                    (3, '終了');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_status_types WHERE user_status_type_id IN (1, 2, 3);
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS user_statuses;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS user_status_types;
-- +goose StatementEnd
