-- Usersテーブルの作成
CREATE TABLE Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Attendance_typesテーブルの作成
CREATE TABLE Attendance_types (
    attendance_type_id INT PRIMARY KEY,
    action_type VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Attendancesテーブルの作成
CREATE TABLE Attendances (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    attendance_type INT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    date VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (attendance_type) REFERENCES Attendance_types(attendance_type_id)
);

-- Attendance_typesテーブルにデータをINSERT
INSERT INTO Attendance_types (attendance_type_id, action_type) VALUES
(1, '出勤退勤'),
(2, '休憩'),
(3, 'お祈り');
