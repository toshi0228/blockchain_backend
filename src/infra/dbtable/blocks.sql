CREATE TABLE IF NOT EXISTS %s
(
    id BIGINT UNSIGNED AUTO_INCREMENT,
    nonce INT UNSIGNED,
    previous_hash VARCHAR(255),
    transactions LONGTEXT,
    timestamp BIGINT UNSIGNED,
    hash VARCHAR(255),
    PRIMARY KEY (id)
)
