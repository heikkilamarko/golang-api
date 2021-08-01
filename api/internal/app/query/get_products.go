package query

import (
	"context"
	"product-api/internal/domain"
)

type GetProducts struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetProductsHandler struct {
	r domain.ProductRepository
}

func NewGetProductsHandler(r domain.ProductRepository) *GetProductsHandler {
	return &GetProductsHandler{r}
}

func (h *GetProductsHandler) Handle(ctx context.Context, q *GetProducts) ([]*domain.Product, error) {
	return h.r.GetProducts(ctx, &domain.GetProductsQuery{
		Offset: q.Offset,
		Limit:  q.Limit,
	})
}
