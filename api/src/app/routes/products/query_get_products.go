package products

import (
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// GetProducts query
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	query, verr := parseGetProductsRequest(r)

	if verr != nil {
		goutils.WriteBadRequest(w, verr.ValidationErrors)
		return
	}

	products, err := c.Repository.GetProducts(r.Context(), query)

	if err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, products, query)
}

func parseGetProductsRequest(r *http.Request) (*GetProductsQuery, *goutils.ValidationError) {
	validationErrors := map[string]string{}

	var offset int = 0
	var limit int = constants.PaginationLimitMax

	var err error = nil

	if value := goutils.GetRequestFormValueString(r, constants.FieldPaginationOffset); value != "" {
		offset, err = goutils.ParseInt(value)
		if err != nil || offset < 0 {
			validationErrors[constants.FieldPaginationOffset] = constants.ErrCodeInvalidOffset
		}
	}

	if value := goutils.GetRequestFormValueString(r, constants.FieldPaginationLimit); value != "" {
		limit, err = goutils.ParseInt(value)
		if err != nil || limit < 1 || constants.PaginationLimitMax < limit {
			validationErrors[constants.FieldPaginationLimit] = constants.ErrCodeInvalidLimit
		}
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &GetProductsQuery{offset, limit}, nil
}
