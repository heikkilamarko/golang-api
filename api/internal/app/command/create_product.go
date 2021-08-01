package command

import (
	"context"
	"product-api/internal/domain"
)

type CreateProduct struct {
	Product *domain.Product `json:"todo"`
}

type CreateProductHandler struct {
	r domain.ProductRepository
}

func NewCreateProductHandler(r domain.ProductRepository) *CreateProductHandler {
	return &CreateProductHandler{r}
}

func (h *CreateProductHandler) Handle(ctx context.Context, c *CreateProduct) error {
	return h.r.CreateProduct(ctx, c.Product)
}
