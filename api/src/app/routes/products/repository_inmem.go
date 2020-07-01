// +build data_inmem

package products

import (
	"context"
	"products-api/app/utils"
	"sort"
)

type repository struct {
	id       int
	products map[int]*product
}

// Initialize method
func (repo *repository) Initialize() {
	repo.products = map[int]*product{}

	repo.id = 1

	repo.products[1] = &product{
		ID:          repo.id,
		Name:        "Inmem product 1",
		Description: "Inmem description 1",
		Price:       100,
		Comment:     "Inmem comment 1",
	}

	repo.id++

	repo.products[2] = &product{
		ID:          repo.id,
		Name:        "Inmem product 2",
		Description: "Inmem description 2",
		Price:       200,
		Comment:     "Inmem comment 2",
	}

	repo.id++
}

func (repo *repository) getProducts(ctx context.Context, query *getProductsQuery) ([]*product, error) {
	products := []*product{}

	for _, product := range repo.products {
		products = append(products, product)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})

	return products, nil
}

func (repo *repository) getProduct(ctx context.Context, id int) (*product, error) {

	product, ok := repo.products[id]

	if !ok {
		return nil, utils.ErrNotFound
	}

	return product, nil
}

func (repo *repository) createProduct(ctx context.Context, p *product) error {
	p.ID = repo.id
	repo.id++
	repo.products[p.ID] = p
	return nil
}

func (repo *repository) updateProduct(ctx context.Context, p *product) error {
	product, ok := repo.products[p.ID]

	if !ok {
		return utils.ErrNotFound
	}

	*product = *p

	return nil
}

func (repo *repository) deleteProduct(ctx context.Context, id int) error {
	_, ok := repo.products[id]

	if !ok {
		return utils.ErrNotFound
	}

	delete(repo.products, id)

	return nil
}
