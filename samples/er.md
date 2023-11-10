```mermaid
erDiagram
    users ||--o{ attendances : ""
    attendances ||--|| attendance_types : ""
    users ||--o{ user_statuses : ""
    user_statuses ||--|| user_status_types : ""
    users ||--o{ current_activity : ""
    attendances ||--o{ current_activity : ""

    users {
        int id PK
        string name
        string email
        string firebase_id
        timestamp created_at
        timestamp updated_at
    }

    attendances {
        int id PK
        int user_id FK
        int attendance_type FK
        timestamp start_time "フロント側から受け取った開始時間"
        timestamp end_time "フロント側から受け取った終了時間"
        string date "YYYY-MM-DD形式の日付"
        int year "年を表す"
        timestamp created_at "レコードの生成と同時に生まれるカラム"
        timestamp updated_at "レコードの生成と同時に生まれるカラム"
    }

    attendance_types {
        int attendance_type_id PK "1:作業 　2:休憩 　3:お祈り"
        string action_type
        timestamp created_at
        timestamp updated_at
    }

    user_status_types {
        int user_status_type_id PK "1:作業中 　2:休憩中 　3:終了"
        string user_status_type
        timestamp created_at
        timestamp updated_at
    }

    user_statuses {
        int user_id PK,FK "ユーザーID"
        int status_id FK "ユーザーの現在の状態ID"
        timestamp updated_at "最終更新時間"
    }

    current_activity {
        int user_id PK, FK "ユーザーID"
        int work_id FK "作業開始のattendance ID"
        int break_id FK "休憩開始のattendance ID" }

```
