CREATE TABLE IF NOT EXISTS %s
(
    id CHAR(40) NOT NULL,
    senderAddress CHAR(40),
    recipientAddress CHAR(40),
    amount DECIMAL(20,2),
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id)
)
