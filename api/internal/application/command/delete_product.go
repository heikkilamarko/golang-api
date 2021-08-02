package command

import (
	"context"
	"product-api/internal/ports"
)

type DeleteProduct struct {
	ID int `json:"id"`
}

type DeleteProductHandler struct {
	r ports.ProductRepository
}

func NewDeleteProductHandler(r ports.ProductRepository) *DeleteProductHandler {
	return &DeleteProductHandler{r}
}

func (h *DeleteProductHandler) Handle(ctx context.Context, c *DeleteProduct) error {
	return h.r.DeleteProduct(ctx, c.ID)
}
