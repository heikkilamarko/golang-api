package products

import (
	"context"
	"products-api/app/utils"
	"sort"
)

// InMemRepository struct
type InMemRepository struct {
	id       int
	products map[int]*Product
}

// NewInMemRepository func
func NewInMemRepository() *InMemRepository {
	return &InMemRepository{
		id:       1,
		products: map[int]*Product{},
	}
}

// GetProducts method
func (r *InMemRepository) GetProducts(ctx context.Context, query *GetProductsQuery) ([]*Product, error) {
	products := []*Product{}

	for _, product := range r.products {
		products = append(products, product)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})

	return products, nil
}

// GetProduct method
func (r *InMemRepository) GetProduct(ctx context.Context, query *GetProductQuery) (*Product, error) {

	product, ok := r.products[query.ID]

	if !ok {
		return nil, utils.ErrNotFound
	}

	return product, nil
}

// CreateProduct method
func (r *InMemRepository) CreateProduct(ctx context.Context, command *CreateProductCommand) error {
	p := command.Product
	p.ID = r.id
	r.id++
	r.products[p.ID] = p
	return nil
}

// UpdateProduct method
func (r *InMemRepository) UpdateProduct(ctx context.Context, command *UpdateProductCommand) error {
	p := command.Product

	product, ok := r.products[p.ID]

	if !ok {
		return utils.ErrNotFound
	}

	*product = *p

	return nil
}

// DeleteProduct method
func (r *InMemRepository) DeleteProduct(ctx context.Context, command *DeleteProductCommand) error {
	_, ok := r.products[command.ID]

	if !ok {
		return utils.ErrNotFound
	}

	delete(r.products, command.ID)

	return nil
}
