package query

import (
	"context"
	"product-api/internal/domain"
	"product-api/internal/ports"
)

type GetPriceRangeHandler struct {
	r ports.ProductRepository
}

func NewGetPriceRangeHandler(r ports.ProductRepository) *GetPriceRangeHandler {
	return &GetPriceRangeHandler{r}
}

func (h *GetPriceRangeHandler) Handle(ctx context.Context) (*domain.PriceRange, error) {
	return h.r.GetPriceRange(ctx)
}
