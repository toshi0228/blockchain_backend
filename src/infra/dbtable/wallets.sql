CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(40) NOT NULL,
    user_id CHAR(40),
    blockchain_address CHAR(40),
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id)
)
