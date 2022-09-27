CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(32) NOT NULL,
    nonce INT UNSIGNED,
    previous_hash CHAR(32),
    transactions_hash CHAR(32),
    timestamp INT UNSIGNED,
    PRIMARY KEY (id)
)
