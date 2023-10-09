-- +goose Up

-- +goose StatementBegin
CREATE TABLE User (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE Activities (
    activity_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    activity_date DATE,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES User(user_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE Play (
    play_id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id INT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES Activities(activity_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE StudyStartEnd (
    study_id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id INT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES Activities(activity_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE BreakStartEnd (
    break_id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id INT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES Activities(activity_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE BreakStartEnd;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE StudyStartEnd;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE Play;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE Activities;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE User;
-- +goose StatementEnd
