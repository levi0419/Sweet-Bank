-- name: CreateEntry :one
INSERT INTO entries (
  amount 
) VALUES (
  $1
  
) RETURNING *;


-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntry :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateEntry :one
UPDATE entries
set amount = $2
WHERE id = $1
RETURNING *;


-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;
