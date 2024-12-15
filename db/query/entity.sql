-- name: CreateEntity :one
INSERT INTO
    entries (account_id, amount)
VALUES ($1, $2)
RETURNING
    *;

-- name: GetEntity :one
SELECT * FROM entries WHERE id = $1 LIMIT 1;

-- name: ListEntities :many
SELECT *
FROM entries
WHERE
    account_id = $1
ORDER BY id
LIMIT $2
OFFSET
    $3;