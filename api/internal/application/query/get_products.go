package query

import (
	"context"
	"product-api/internal/domain"
	"product-api/internal/ports"
)

type GetProducts struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetProductsHandler struct {
	r ports.ProductRepository
}

func NewGetProductsHandler(r ports.ProductRepository) *GetProductsHandler {
	return &GetProductsHandler{r}
}

func (h *GetProductsHandler) Handle(ctx context.Context, q *GetProducts) ([]*domain.Product, error) {
	return h.r.GetProducts(ctx, &ports.GetProductsQuery{
		Offset: q.Offset,
		Limit:  q.Limit,
	})
}
