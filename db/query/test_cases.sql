-- name: GetTestCase :one
SELECT id, name, hidden, input, output, challenge_id, created_at, updated_at
FROM test_cases
WHERE id = $1
AND deleted_at IS NULL;

-- name: ListHiddenTestCases :many
SELECT id, name, hidden, input, output, challenge_id, created_at, updated_at
FROM test_cases
WHERE challenge_id = $1
AND hidden IS TRUE
AND deleted_at IS NULL
LIMIT $2
OFFSET $3;

-- name: ListExposedTestCases :many
SELECT id, name, hidden, input, output, challenge_id, created_at, updated_at
FROM test_cases
WHERE challenge_id = $1
AND hidden IS FALSE
AND deleted_at IS NULL
LIMIT $2
OFFSET $3;

-- name: ListTestCases :many
SELECT id, name, hidden, input, output, challenge_id, created_at, updated_at
FROM test_cases
WHERE challenge_id = $1
AND deleted_at IS NULL
LIMIT $2
OFFSET $3;

-- name: CreateTestCase :one
INSERT INTO test_cases (name, hidden, input, output, challenge_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, hidden, input, output, challenge_id, created_at, updated_at;

-- name: UpdateTestCase :one
UPDATE test_cases SET name = $2, hidden = $3, input = $4, output = $5, updated_at = NOW() WHERE id = $1 AND deleted_at IS NULL RETURNING id, name, hidden, input, output, challenge_id, created_at, updated_at;

-- name: DeleteTestCase :exec
UPDATE test_cases SET deleted_at = NOW() WHERE id = $1;

-- name: CountTestCases :one
SELECT COUNT(*)
FROM test_cases
WHERE deleted_at IS NULL;
