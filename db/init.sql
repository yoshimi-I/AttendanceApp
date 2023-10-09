-- Userテーブルの作成
CREATE TABLE User (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Activitiesテーブルの作成
CREATE TABLE Activities (
    activity_id CHAR(36) DEFAULT (UUID()) PRIMARY KEY,
    user_id INT,
    activity_date DATE,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES User(user_id)
);

-- Playテーブルの作成
CREATE TABLE Play (
    play_id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id CHAR(36),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES Activities(activity_id)
);

-- StudyStartEndテーブルの作成
CREATE TABLE StudyStartEnd (
    study_id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id CHAR(36),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES Activities(activity_id)
);

-- BreakStartEndテーブルの作成
CREATE TABLE BreakStartEnd (
    break_id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id CHAR(36),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES Activities(activity_id)
);
