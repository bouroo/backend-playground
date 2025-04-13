package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v6"
)

type Item struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	CategoryID uuid.UUID `json:"category"`
	Price      float64   `json:"price"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  null.Time `json:"deleted_at"`
}
