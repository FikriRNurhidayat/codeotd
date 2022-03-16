package entity

import (
	"time"

	"github.com/google/uuid"
)

type Challenge struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Body        string     `json:"body,omitempty"`
	Description string     `json:"description"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
