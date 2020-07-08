// +build data_inmem

package products

import (
	"context"
	"products-api/app/utils"
	"sort"
)

type repository struct {
	id       int
	products map[int]*Product
}

func (repo *repository) initialize() {
	repo.products = map[int]*Product{}

	repo.id = 1

	repo.products[1] = &Product{
		ID:          repo.id,
		Name:        "Inmem product 1",
		Description: "Inmem description 1",
		Price:       100,
		Comment:     "Inmem comment 1",
	}

	repo.id++

	repo.products[2] = &Product{
		ID:          repo.id,
		Name:        "Inmem product 2",
		Description: "Inmem description 2",
		Price:       200,
		Comment:     "Inmem comment 2",
	}

	repo.id++
}

func (repo *repository) GetProducts(ctx context.Context, query *GetProductsQuery) ([]*Product, error) {
	products := []*Product{}

	for _, product := range repo.products {
		products = append(products, product)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})

	return products, nil
}

func (repo *repository) GetProduct(ctx context.Context, query *GetProductQuery) (*Product, error) {

	product, ok := repo.products[query.ID]

	if !ok {
		return nil, utils.ErrNotFound
	}

	return product, nil
}

func (repo *repository) CreateProduct(ctx context.Context, command *CreateProductCommand) error {
	p := command.Product
	p.ID = repo.id
	repo.id++
	repo.products[p.ID] = p
	return nil
}

func (repo *repository) UpdateProduct(ctx context.Context, command *UpdateProductCommand) error {
	p := command.Product

	product, ok := repo.products[p.ID]

	if !ok {
		return utils.ErrNotFound
	}

	*product = *p

	return nil
}

func (repo *repository) DeleteProduct(ctx context.Context, command *DeleteProductCommand) error {
	_, ok := repo.products[command.ID]

	if !ok {
		return utils.ErrNotFound
	}

	delete(repo.products, command.ID)

	return nil
}
