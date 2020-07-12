package products

import (
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// GetProducts query
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	p := newGetProductsRequestParser(r).parse()

	if !p.IsValid() {
		goutils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	products, err := c.Repository.GetProducts(r.Context(), p.query)

	if err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, products, p.query)
}

func newGetProductsRequestParser(r *http.Request) *getProductsRequestParser {
	return &getProductsRequestParser{goutils.RequestValidator{Request: r}, nil}
}

type getProductsRequestParser struct {
	goutils.RequestValidator
	query *GetProductsQuery
}

func (p *getProductsRequestParser) parse() *getProductsRequestParser {
	validationErrors := map[string]string{}

	var offset int = 0
	var limit int = constants.PaginationLimitMax

	var err error = nil

	if value := goutils.GetRequestFormValueString(p.Request, constants.FieldPaginationOffset); value != "" {
		offset, err = goutils.ParseInt(value)
		if err != nil || offset < 0 {
			validationErrors[constants.FieldPaginationOffset] = constants.ErrCodeInvalidOffset
		}
	}

	if value := goutils.GetRequestFormValueString(p.Request, constants.FieldPaginationLimit); value != "" {
		limit, err = goutils.ParseInt(value)
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
