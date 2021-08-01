package command

import (
	"context"
	"product-api/internal/domain"
)

type DeleteProduct struct {
	ID int `json:"id"`
}

type DeleteProductHandler struct {
	r domain.ProductRepository
}

func NewDeleteProductHandler(r domain.ProductRepository) *DeleteProductHandler {
	return &DeleteProductHandler{r}
}

func (h *DeleteProductHandler) Handle(ctx context.Context, c *DeleteProduct) error {
	return h.r.DeleteProduct(ctx, c.ID)
}
