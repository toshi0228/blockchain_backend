CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(40) NOT NULL,
    sender_address CHAR(40),
    recipient_address CHAR(40),
    amount INT UNSIGNED,
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id)
)
