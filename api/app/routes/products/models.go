package products

import (
	"time"
)

type product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Price       float64    `json:"price"`
	Comment     *string    `json:"comment,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type priceRange struct {
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
}

type getProductsQuery struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type getProductQuery struct {
	ID int `json:"id"`
}

type createProductCommand struct {
	Product *product `json:"product"`
}

type updateProductCommand struct {
	Product *product `json:"product"`
}

type deleteProductCommand struct {
	ID int `json:"id"`
}
