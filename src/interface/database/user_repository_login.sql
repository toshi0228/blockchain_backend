SELECT
    users.id,
    users.name,
    users.password,
    users.created_at,
    users.updated_at,
    w.id AS wallet_id,
    w.user_id,
    w.blockchain_address,
    w.created_at AS wallet_created_at,
    w.updated_at AS wallet_updated_at
FROM
    users LEFT JOIN wallets w on users.id = w.user_id
WHERE
    name = ? && password = ?

