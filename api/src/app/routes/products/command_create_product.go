package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// CreateProduct command
func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	p := newCreateProductRequestParser(r)
	p.parse()

	if !p.IsValid() {
		utils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	if err := c.Repository.CreateProduct(r.Context(), p.command); err != nil {
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteCreated(w, p.command.Product, nil)
}

func newCreateProductRequestParser(r *http.Request) *createProductRequestParser {
	return &createProductRequestParser{utils.RequestValidator{Request: r}, nil}
}

type createProductRequestParser struct {
	utils.RequestValidator
	command *CreateProductCommand
}

func (v *createProductRequestParser) parse() {
	validationErrors := map[string]string{}

	product := &Product{}
	if err := json.NewDecoder(v.Request.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		v.ValidationErrors = validationErrors
	} else {
		v.command = &CreateProductCommand{product}
	}
}
