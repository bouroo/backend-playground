package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v6"
)

type DiscountType string

const (
	DiscountTypePercentage DiscountType = "percentage"
	DiscountTypeFixed      DiscountType = "fixed"
)

type Discount struct {
	Type  DiscountType `json:"type"`
	Value float64      `json:"value"`
}

type Campaign struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Discount    Discount  `json:"discount"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   null.Time `json:"deleted_at"`
}
