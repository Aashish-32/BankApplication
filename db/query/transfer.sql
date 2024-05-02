-- name: CreateTransfer :one
INSERT into transfers
(from_account_id, to_account_id,amount)
VALUES($1,$2,$3)
RETURNING *;


-- name: GetTransfer :one
SELECT * from transfers where id=$1 LIMIT 1;

-- name: ListTransfer :many
SELECT * from transfers
WHERE from_account_id=$1 AND to_account_id=$2
ORDER BY id 
LIMIT $3
OFFSET $4;