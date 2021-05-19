package products

import (
	_ "embed"
	"time"
)

var (
	//go:embed sql/get_products.sql
	qetProductsSQL string
	//go:embed sql/get_product.sql
	qetProductSQL string
	//go:embed sql/create_product.sql
	createProductSQL string
	//go:embed sql/update_product.sql
	updateProductSQL string
	//go:embed sql/delete_product.sql
	deleteProductSQL string
	//go:embed sql/get_price_range.sql
	getPriceRangeSQL string
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
