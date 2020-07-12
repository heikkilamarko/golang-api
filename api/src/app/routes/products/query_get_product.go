package products

import (
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// GetProduct query
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	p := newGetProductRequestParser(r).parse()

	if !p.IsValid() {
		goutils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	product, err := c.Repository.GetProduct(r.Context(), p.query)

	if err != nil {
		switch err {
		case goutils.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, product, nil)
}

func newGetProductRequestParser(r *http.Request) *getProductRequestParser {
	return &getProductRequestParser{goutils.RequestValidator{Request: r}, nil}
}

type getProductRequestParser struct {
	goutils.RequestValidator
	query *GetProductQuery
}

func (p *getProductRequestParser) parse() *getProductRequestParser {
	validationErrors := map[string]string{}

	id, err := goutils.GetRequestVarInt(p.Request, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		p.ValidationErrors = validationErrors
	} else {
		p.query = &GetProductQuery{id}
	}

	return p
}
