-- Usersテーブル
CREATE TABLE users (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       user_key VARCHAR(255) UNIQUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 作業の種類を含めるテーブル
CREATE TABLE attendance_types (
                                  attendance_type_id INT PRIMARY KEY,
                                  action_type VARCHAR(100) NOT NULL,
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Userの現在の状態の種類を管理するテーブル
CREATE TABLE user_status_types (
                                   user_status_type_id INT PRIMARY KEY,
                                   user_status_type VARCHAR(100) NOT NULL,
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ユーザーの作業履歴を見るテーブル
CREATE TABLE attendances (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             user_id INT,
                             attendance_type INT,
                             start_time TIMESTAMP,
                             end_time TIMESTAMP,
                             date VARCHAR(15),
                             year INT,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             FOREIGN KEY (user_id) REFERENCES users(id),
                             FOREIGN KEY (attendance_type) REFERENCES attendance_types(attendance_type_id)
);

-- Userの現在の状態を管理するテーブル
CREATE TABLE user_statuses (
                               user_id INT PRIMARY KEY,
                               status_id INT,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               FOREIGN KEY (user_id) REFERENCES users(id),
                               FOREIGN KEY (status_id) REFERENCES user_status_types(user_status_type_id)
);

-- Userの直近の行動のpkを保持するテーブル
CREATE TABLE current_activities (
                                    user_id INT PRIMARY KEY,
                                    work_id INT NULL,  -- NULL 許容
                                    break_id INT NULL, -- NULL 許容
                                    FOREIGN KEY (user_id) REFERENCES users(id),
                                    FOREIGN KEY (work_id) REFERENCES attendances(id),
                                    FOREIGN KEY (break_id) REFERENCES attendances(id)
);

-- Attendance_typesテーブルにデータをINSERT
INSERT INTO attendance_types (attendance_type_id, action_type) VALUES
                                                                   (1, '作業'),
                                                                   (2, '休憩'),
                                                                   (3, 'お祈り');

-- user_status_typesテーブルにデータをINSERT
INSERT INTO user_status_types (user_status_type_id, user_status_type) VALUES
                                                                          (1, '作業中'),
                                                                          (2, '休憩中'),
                                                                          (3, '終了');
