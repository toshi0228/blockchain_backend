CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(32) NOT NULL,
    nonce INT UNSIGNED,
    previous_hash VARCHAR(255),
    transactions LONGTEXT,
    timestamp BIGINT UNSIGNED,
    hash VARCHAR(255),
    PRIMARY KEY (id)
)
