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
                             time TIMESTAMP,
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
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               FOREIGN KEY (user_id) REFERENCES users(id),
                               FOREIGN KEY (status_id) REFERENCES user_status_types(user_status_type_id)
);


-- Attendance_typesテーブルにデータをINSERT
INSERT INTO attendance_types (attendance_type_id, action_type) VALUES
                                                                   (1, '作業開始'),
                                                                   (2, '作業終了'),
                                                                   (3, '休憩開始'),
                                                                   (4, '休憩終了'),
                                                                   (5, 'お祈り');

-- user_status_typesテーブルにデータをINSERT
INSERT INTO user_status_types (user_status_type_id, user_status_type) VALUES
                                                                          (1, '作業中'),
                                                                          (2, '休憩中'),
                                                                          (3, '終了');
