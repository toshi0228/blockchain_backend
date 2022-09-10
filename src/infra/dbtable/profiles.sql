CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(40) NOT NULL,
    user_id CHAR(40),
    bio TEXT,
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)
