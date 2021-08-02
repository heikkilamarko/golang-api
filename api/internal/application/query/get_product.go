package query

import (
	"context"
	"product-api/internal/domain"
	"product-api/internal/ports"
)

type GetProduct struct {
	ID int `json:"id"`
}

type GetProductHandler struct {
	r ports.ProductRepository
}

func NewGetProductHandler(r ports.ProductRepository) *GetProductHandler {
	return &GetProductHandler{r}
}

func (h *GetProductHandler) Handle(ctx context.Context, q *GetProduct) (*domain.Product, error) {
	return h.r.GetProduct(ctx, q.ID)
}
