// Code generated by sqlc. DO NOT EDIT.
// source: challenges.sql

package dao

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const countChallenges = `-- name: CountChallenges :one
SELECT COUNT(*) FROM challenges
WHERE deleted_at IS NULL
`

func (q *Queries) CountChallenges(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countChallenges)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createChallenge = `-- name: CreateChallenge :one
INSERT INTO challenges (title, description, body)
VALUES ($1, $2, $3)
RETURNING id, title, description, body, created_at, updated_at
`

type CreateChallengeParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

type CreateChallengeRow struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) CreateChallenge(ctx context.Context, arg CreateChallengeParams) (CreateChallengeRow, error) {
	row := q.db.QueryRowContext(ctx, createChallenge, arg.Title, arg.Description, arg.Body)
	var i CreateChallengeRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteChallenge = `-- name: DeleteChallenge :exec
UPDATE challenges
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteChallenge(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteChallenge, id)
	return err
}

const getChallenge = `-- name: GetChallenge :one
SELECT id, title, description, body, created_at, updated_at
FROM challenges
WHERE id = $1 AND deleted_at IS NULL
`

type GetChallengeRow struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) GetChallenge(ctx context.Context, id uuid.UUID) (GetChallengeRow, error) {
	row := q.db.QueryRowContext(ctx, getChallenge, id)
	var i GetChallengeRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listChallenges = `-- name: ListChallenges :many
SELECT id, title, description, created_at, updated_at
FROM challenges
WHERE deleted_at IS NULL
LIMIT $1
OFFSET $2
`

type ListChallengesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListChallengesRow struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) ListChallenges(ctx context.Context, arg ListChallengesParams) ([]ListChallengesRow, error) {
	rows, err := q.db.QueryContext(ctx, listChallenges, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListChallengesRow
	for rows.Next() {
		var i ListChallengesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateChallenge = `-- name: UpdateChallenge :one
UPDATE challenges
SET title = $2,
    description = $3,
    body  = $4
WHERE id = $1
AND deleted_at IS NULL
RETURNING id, title, description, body, created_at, updated_at
`

type UpdateChallengeParams struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
}

type UpdateChallengeRow struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) UpdateChallenge(ctx context.Context, arg UpdateChallengeParams) (UpdateChallengeRow, error) {
	row := q.db.QueryRowContext(ctx, updateChallenge,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Body,
	)
	var i UpdateChallengeRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
