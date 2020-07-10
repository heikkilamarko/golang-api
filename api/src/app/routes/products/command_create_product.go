package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// CreateProduct command
func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	v := newCreateProductCommandParser(r)
	v.parse()

	if !v.IsValid() {
		utils.WriteBadRequest(w, v.ValidationErrors)
		return
	}

	if err := c.Repository.CreateProduct(r.Context(), v.command); err != nil {
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteCreated(w, v.command.Product, nil)
}

func newCreateProductCommandParser(r *http.Request) *createProductCommandParser {
	return &createProductCommandParser{utils.RequestValidator{Request: r}, nil}
}

type createProductCommandParser struct {
	utils.RequestValidator
	command *CreateProductCommand
}

func (v *createProductCommandParser) parse() {
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
