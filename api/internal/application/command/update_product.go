package command

import (
	"context"
	"product-api/internal/domain"
	"product-api/internal/ports"
)

type UpdateProduct struct {
	Product *domain.Product `json:"todo"`
}

type UpdateProductHandler struct {
	r ports.ProductRepository
}

func NewUpdateProductHandler(r ports.ProductRepository) *UpdateProductHandler {
	return &UpdateProductHandler{r}
}

func (h *UpdateProductHandler) Handle(ctx context.Context, c *UpdateProduct) error {
	return h.r.UpdateProduct(ctx, c.Product)
}
