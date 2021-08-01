package command

import (
	"context"
	"product-api/internal/domain"
)

type UpdateProduct struct {
	Product *domain.Product `json:"todo"`
}

type UpdateProductHandler struct {
	r domain.ProductRepository
}

func NewUpdateProductHandler(r domain.ProductRepository) *UpdateProductHandler {
	return &UpdateProductHandler{r}
}

func (h *UpdateProductHandler) Handle(ctx context.Context, c *UpdateProduct) error {
	return h.r.UpdateProduct(ctx, c.Product)
}
