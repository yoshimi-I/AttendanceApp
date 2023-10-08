-- Userテーブルの作成
CREATE TABLE User (
                      user_id INT PRIMARY KEY,
                      name VARCHAR(255),
                      email VARCHAR(255) UNIQUE,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Attendancesテーブルの作成
CREATE TABLE Attendances (
                             attendance_id INT PRIMARY KEY,
                             user_id INT,
                             attendance_date DATE,
                             clock_in_time TIMESTAMP,
                             clock_out_time TIMESTAMP,
                             notes TEXT,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             FOREIGN KEY (user_id) REFERENCES User(user_id)
);

-- Breaksテーブルの作成
CREATE TABLE Breaks (
                        break_id INT PRIMARY KEY,
                        attendance_id INT,
                        start_time TIMESTAMP,
                        end_time TIMESTAMP,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        FOREIGN KEY (attendance_id) REFERENCES Attendances(attendance_id)
);

-- PrayerTimesテーブルの作成
CREATE TABLE PrayerTimes (
                             prayer_id INT PRIMARY KEY,
                             attendance_id INT,
                             fajr TIMESTAMP,
                             dhuhr TIMESTAMP,
                             asr TIMESTAMP,
                             maghrib TIMESTAMP,
                             isha TIMESTAMP,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             FOREIGN KEY (attendance_id) REFERENCES Attendances(attendance_id)
);
