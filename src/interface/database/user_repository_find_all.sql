SELECT
    *
FROM
    users
LEFT JOIN wallets w on users.id = w.user_id
