package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// CreateProduct command
func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	p := newCreateProductRequestParser(r).parse()

	if !p.IsValid() {
		goutils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	if err := c.Repository.CreateProduct(r.Context(), p.command); err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteCreated(w, p.command.Product, nil)
}

func newCreateProductRequestParser(r *http.Request) *createProductRequestParser {
	return &createProductRequestParser{goutils.RequestValidator{Request: r}, nil}
}

type createProductRequestParser struct {
	goutils.RequestValidator
	command *CreateProductCommand
}

func (p *createProductRequestParser) parse() *createProductRequestParser {
	validationErrors := map[string]string{}

	product := &Product{}
	if err := json.NewDecoder(p.Request.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		p.ValidationErrors = validationErrors
	} else {
		p.command = &CreateProductCommand{product}
	}

	return p
}
