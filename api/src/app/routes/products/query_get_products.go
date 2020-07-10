package products

import (
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// GetProducts query
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	p := newGetProductsRequestParser(r).parse()

	if !p.IsValid() {
		utils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	products, err := c.Repository.GetProducts(r.Context(), p.query)

	if err != nil {
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteOK(w, products, p.query)
}

func newGetProductsRequestParser(r *http.Request) *getProductsRequestParser {
	return &getProductsRequestParser{utils.RequestValidator{Request: r}, nil}
}

type getProductsRequestParser struct {
	utils.RequestValidator
	query *GetProductsQuery
}

func (p *getProductsRequestParser) parse() *getProductsRequestParser {
	validationErrors := map[string]string{}

	var offset int = 0
	var limit int = constants.PaginationLimitMax

	var err error = nil

	if value := utils.GetRequestFormValueString(p.Request, constants.FieldPaginationOffset); value != "" {
		offset, err = utils.ParseInt(value)
		if err != nil || offset < 0 {
			validationErrors[constants.FieldPaginationOffset] = constants.ErrCodeInvalidOffset
		}
	}

	if value := utils.GetRequestFormValueString(p.Request, constants.FieldPaginationLimit); value != "" {
		limit, err = utils.ParseInt(value)
		if err != nil || limit < 1 || constants.PaginationLimitMax < limit {
			validationErrors[constants.FieldPaginationLimit] = constants.ErrCodeInvalidLimit
		}
	}

	if 0 < len(validationErrors) {
		p.ValidationErrors = validationErrors
	} else {
		p.query = &GetProductsQuery{offset, limit}
	}

	return p
}
