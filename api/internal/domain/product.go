package domain

import "time"

type Product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Price       float64    `json:"price"`
	Comment     *string    `json:"comment,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type PriceRange struct {
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
}
