package query

import (
	"context"
	"product-api/internal/domain"
)

type GetProduct struct {
	ID int `json:"id"`
}

type GetProductHandler struct {
	r domain.ProductRepository
}

func NewGetProductHandler(r domain.ProductRepository) *GetProductHandler {
	return &GetProductHandler{r}
}

func (h *GetProductHandler) Handle(ctx context.Context, q *GetProduct) (*domain.Product, error) {
	return h.r.GetProduct(ctx, q.ID)
}
