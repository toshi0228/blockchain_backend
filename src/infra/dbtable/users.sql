CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(40) NOT NULL,
    name VARCHAR(255),
    password VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id)
)
