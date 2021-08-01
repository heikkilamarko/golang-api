package query

import (
	"context"
	"product-api/internal/domain"
)

type GetPriceRangeHandler struct {
	r domain.ProductRepository
}

func NewGetPriceRangeHandler(r domain.ProductRepository) *GetPriceRangeHandler {
	return &GetPriceRangeHandler{r}
}

func (h *GetPriceRangeHandler) Handle(ctx context.Context) (*domain.PriceRange, error) {
	return h.r.GetPriceRange(ctx)
}
