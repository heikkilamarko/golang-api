package products

import (
	"net/http"
	"products-api/app/utils"
	"strconv"

	"github.com/heikkilamarko/goutils"
)

// GetProducts query
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetProductsRequest(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	products, err := c.repository.getProducts(r.Context(), query)

	if err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, products, query)
}

func parseGetProductsRequest(r *http.Request) (*getProductsQuery, error) {
	errorMap := map[string]string{}

	offset := 0
	limit := utils.LimitMaxPageSize

	var err error

	if value := r.FormValue(utils.FieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			errorMap[utils.FieldPaginationOffset] = utils.ErrCodeInvalidOffset
		}
	}

	if value := r.FormValue(utils.FieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || utils.LimitMaxPageSize < limit {
			errorMap[utils.FieldPaginationLimit] = utils.ErrCodeInvalidLimit
		}
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &getProductsQuery{offset, limit}, nil
}
