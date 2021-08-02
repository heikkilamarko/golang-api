package ports

import (
	"context"
	"product-api/internal/domain"
)

type GetProductsQuery struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ProductRepository interface {
	GetProducts(ctx context.Context, query *GetProductsQuery) ([]*domain.Product, error)
	GetProduct(ctx context.Context, id int) (*domain.Product, error)
	CreateProduct(ctx context.Context, p *domain.Product) error
	UpdateProduct(ctx context.Context, p *domain.Product) error
	DeleteProduct(ctx context.Context, id int) error
	GetPriceRange(ctx context.Context) (*domain.PriceRange, error)
}
