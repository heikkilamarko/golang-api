package command

import (
	"context"
	"product-api/internal/domain"
	"product-api/internal/ports"
)

type CreateProduct struct {
	Product *domain.Product
}

type CreateProductHandler struct {
	r ports.ProductRepository
}

func NewCreateProductHandler(r ports.ProductRepository) *CreateProductHandler {
	return &CreateProductHandler{r}
}

func (h *CreateProductHandler) Handle(ctx context.Context, c *CreateProduct) error {
	c.Product.SetCreateTimestamps()
	return h.r.CreateProduct(ctx, c.Product)
}
