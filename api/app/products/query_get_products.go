package products

import (
	"net/http"
	"products-api/app/utils"
	"strconv"
)

// GetProducts query
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetProductsRequest(r)

	if err != nil {
		c.logError(err)
		utils.WriteValidationError(w, err)
		return
	}

	products, err := c.repository.getProducts(r.Context(), query)

	if err != nil {
		c.logError(err)
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteOK(w, products, query)
}

func parseGetProductsRequest(r *http.Request) (*getProductsQuery, error) {
	errorMap := map[string]string{}

	var offset int = 0
	var limit int = utils.LimitMaxPageSize

	var err error = nil

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
		return nil, utils.NewValidationError(errorMap)
	}

	return &getProductsQuery{offset, limit}, nil
}
