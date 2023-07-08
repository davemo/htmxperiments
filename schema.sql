-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    password TEXT,
    email TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Create securities table
CREATE TABLE IF NOT EXISTS securities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    user_id INTEGER,
    amount INTEGER,
    type TEXT,
    price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Create exercise_requests table
CREATE TABLE IF NOT EXISTS exercise_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    security_id INTEGER,
    amount INTEGER,
    exercise_date DATE,
    status TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (security_id) REFERENCES securities(id)
);