CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    password TEXT
);

CREATE TABLE IF NOT EXISTS securities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS exercise_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    security_id INTEGER,
    amount INTEGER,
    FOREIGN KEY (security_id) REFERENCES securities(id)
);