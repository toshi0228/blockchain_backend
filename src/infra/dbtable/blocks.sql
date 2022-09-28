CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(32) NOT NULL,
    nonce INT UNSIGNED,
    previous_hash VARCHAR(255),
    transactions_hash VARCHAR(255),
    timestamp BIGINT UNSIGNED,
    PRIMARY KEY (id)
)
