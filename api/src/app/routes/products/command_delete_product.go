package products

import (
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// DeleteProduct command
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	p := newDeleteProductRequestParser(r).parse()

	if !p.IsValid() {
		goutils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	if err := c.Repository.DeleteProduct(r.Context(), p.command); err != nil {
		switch err {
		case goutils.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteNoContent(w)
}

func newDeleteProductRequestParser(r *http.Request) *deleteProductRequestParser {
	return &deleteProductRequestParser{goutils.RequestValidator{Request: r}, nil}
}

type deleteProductRequestParser struct {
	goutils.RequestValidator
	command *DeleteProductCommand
}

func (p *deleteProductRequestParser) parse() *deleteProductRequestParser {
	validationErrors := map[string]string{}

	id, err := goutils.GetRequestVarInt(p.Request, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		p.ValidationErrors = validationErrors
	} else {
		p.command = &DeleteProductCommand{id}
	}

	return p
}
