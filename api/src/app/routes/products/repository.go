package products

import "context"

// Repository interface
type Repository interface {
	GetProducts(context.Context, *GetProductsQuery) ([]*Product, error)
	GetProduct(context.Context, *GetProductQuery) (*Product, error)
	CreateProduct(context.Context, *CreateProductCommand) error
	UpdateProduct(context.Context, *UpdateProductCommand) error
	DeleteProduct(context.Context, *DeleteProductCommand) error
}

// NewRepository func
func NewRepository() Repository {
	r := &repository{}
	r.initialize()
	return r
}
