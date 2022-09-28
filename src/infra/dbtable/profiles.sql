CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(32) NOT NULL,
    user_id CHAR(32),
    bio TEXT,
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)
