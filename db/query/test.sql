-- name: GetTest :one
SELECT * FROM tests WHERE id = $1 AND deleted_at IS NULL;

-- name: ListTests :many
SELECT * FROM tests WHERE deleted_at IS NULL;

-- name: CreateTest :one
INSERT INTO tests (title, body) VALUES ($1, $2) RETURNING *;

-- name: UpdateTest :one
UPDATE tests SET title = $2, body = $3 WHERE id = $1 AND deleted_at IS NULL RETURNING *;

-- name: DeleteTest :exec
UPDATE tests SET deleted_at = NOW() WHERE id = $1;
