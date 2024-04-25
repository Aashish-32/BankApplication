-- name: CreateEntry :one

INSERT into entries(account_id,amount) VALUES ($1,$2)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries WHERE id=$1 LIMIT 1;

-- name: ListEntry :many
SELECT * from entries
WHERE account_id=$1
ORDER BY id
LIMIT $2
OFFSET $3;
