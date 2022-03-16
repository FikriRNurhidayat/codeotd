-- name: GetChallenge :one
SELECT id, title, description, body, created_at, updated_at
FROM challenges
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListChallenges :many
SELECT id, title, description, created_at, updated_at
FROM challenges
WHERE deleted_at IS NULL
LIMIT $1
OFFSET $2;

-- name: CountChallenges :one
SELECT COUNT(*) FROM challenges
WHERE deleted_at IS NULL;

-- name: CreateChallenge :one
INSERT INTO challenges (title, description, body)
VALUES ($1, $2, $3)
RETURNING id, title, description, body, created_at, updated_at;

-- name: UpdateChallenge :one
UPDATE challenges
SET title = $2,
    description = $3,
    body  = $4
WHERE id = $1
AND deleted_at IS NULL
RETURNING id, title, description, body, created_at, updated_at;

-- name: DeleteChallenge :exec
UPDATE challenges
SET deleted_at = NOW()
WHERE id = $1;
