package products

import (
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// GetProducts query
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	v := getProductsQueryValidator{utils.RequestValidator{Request: r}, nil}
	v.parseAndValidate()

	if !v.IsValid() {
		utils.WriteBadRequest(w, v.ValidationErrors)
		return
	}

	products, err := c.Repository.GetProducts(r.Context(), v.query)

	if err != nil {
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteOK(w, products, v.query)
}

type getProductsQueryValidator struct {
	utils.RequestValidator
	query *GetProductsQuery
}

const limitMax = 100

func (v *getProductsQueryValidator) parseAndValidate() {
	validationErrors := map[string]string{}

	var offset int = 0
	var limit int = limitMax

	var err error = nil

	if value := utils.GetRequestFormValueString(v.Request, constants.FieldPaginationOffset); value != "" {
		offset, err = utils.ParseInt(value)
		if err != nil || offset < 0 {
			validationErrors[constants.FieldPaginationOffset] = constants.ErrCodeInvalidOffset
		}
	}

	if value := utils.GetRequestFormValueString(v.Request, constants.FieldPaginationLimit); value != "" {
		limit, err = utils.ParseInt(value)
		if err != nil || limit < 1 || limitMax < limit {
			validationErrors[constants.FieldPaginationLimit] = constants.ErrCodeInvalidLimit
		}
	}

	if 0 < len(validationErrors) {
		v.ValidationErrors = validationErrors
	} else {
		v.query = &GetProductsQuery{offset, limit}
	}
}
