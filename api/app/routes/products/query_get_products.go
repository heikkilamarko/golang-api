package products

import (
	"net/http"
	"products-api/app/constants"
	"strconv"

	"github.com/heikkilamarko/goutils"
)

// GetProducts query
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetProductsRequest(r)

	if err != nil {
		goutils.WriteValidationError(w, err)
		return
	}

	products, err := c.repository.getProducts(r.Context(), query)

	if err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, products, query)
}

func parseGetProductsRequest(r *http.Request) (*getProductsQuery, error) {
	validationErrors := map[string]string{}

	var offset int = 0
	var limit int = constants.PaginationLimitMax

	var err error = nil

	if value := r.FormValue(constants.FieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			validationErrors[constants.FieldPaginationOffset] = constants.ErrCodeInvalidOffset
		}
	}

	if value := r.FormValue(constants.FieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || constants.PaginationLimitMax < limit {
			validationErrors[constants.FieldPaginationLimit] = constants.ErrCodeInvalidLimit
		}
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &getProductsQuery{offset, limit}, nil
}
