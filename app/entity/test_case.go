package entity

import (
	"time"

	"github.com/google/uuid"
)

type TestCase struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Hidden      bool         `json:"hidden"`
	Input       string       `json:"input"`
	Output      string       `json:"output"`
	ChallengeID uuid.UUID    `json:"challenge_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
