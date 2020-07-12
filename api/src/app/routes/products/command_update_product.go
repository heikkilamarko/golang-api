package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// UpdateProduct command
func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p := newUpdateProductRequestParser(r).parse()

	if !p.IsValid() {
		goutils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	if err := c.Repository.UpdateProduct(r.Context(), p.command); err != nil {
		switch err {
		case goutils.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, p.command.Product, nil)
}

func newUpdateProductRequestParser(r *http.Request) *updateProductRequestParser {
	return &updateProductRequestParser{goutils.RequestValidator{Request: r}, nil}
}

type updateProductRequestParser struct {
	goutils.RequestValidator
	command *UpdateProductCommand
}

func (p *updateProductRequestParser) parse() *updateProductRequestParser {
	validationErrors := map[string]string{}

	id, err := goutils.GetRequestVarInt(p.Request, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	product := &Product{}
	if err := json.NewDecoder(p.Request.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		p.ValidationErrors = validationErrors
		return p
	}

	if id != product.ID {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		p.ValidationErrors = validationErrors
	} else {
		p.command = &UpdateProductCommand{product}
	}

	return p
}
