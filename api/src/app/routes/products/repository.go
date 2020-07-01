package products

import "context"

// Repository interface
type Repository interface {
	Initialize()
	getProducts(context.Context, *getProductsQuery) ([]*product, error)
	getProduct(context.Context, int) (*product, error)
	createProduct(context.Context, *product) error
	updateProduct(context.Context, *product) error
	deleteProduct(context.Context, int) error
}

// NewRepository func
func NewRepository() Repository {
	return &repository{}
}
