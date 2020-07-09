package products

import "context"

// Product struct
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Comment     string  `json:"comment"`
}

// GetProductsQuery struct
type GetProductsQuery struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// GetProductQuery struct
type GetProductQuery struct {
	ID int
}

// CreateProductCommand struct
type CreateProductCommand struct {
	Product *Product
}

// UpdateProductCommand struct
type UpdateProductCommand struct {
	Product *Product
}

// DeleteProductCommand struct
type DeleteProductCommand struct {
	ID int
}

// Repository interface
type Repository interface {
	GetProducts(context.Context, *GetProductsQuery) ([]*Product, error)
	GetProduct(context.Context, *GetProductQuery) (*Product, error)
	CreateProduct(context.Context, *CreateProductCommand) error
	UpdateProduct(context.Context, *UpdateProductCommand) error
	DeleteProduct(context.Context, *DeleteProductCommand) error
}
