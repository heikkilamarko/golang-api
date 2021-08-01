package domain

import "context"

type ProductRepository interface {
	GetProducts(ctx context.Context, query *GetProductsQuery) ([]*Product, error)
	GetProduct(ctx context.Context, id int) (*Product, error)
	CreateProduct(ctx context.Context, p *Product) error
	UpdateProduct(ctx context.Context, p *Product) error
	DeleteProduct(ctx context.Context, id int) error
	GetPriceRange(ctx context.Context) (*PriceRange, error)
}
