package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v6"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Hashed    string    `json:"hashed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
