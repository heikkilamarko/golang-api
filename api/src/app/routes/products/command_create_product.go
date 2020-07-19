package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// CreateProduct command
func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	command, verr := parseCreateProductRequest(r)

	if verr != nil {
		goutils.WriteBadRequest(w, verr.ValidationErrors)
		return
	}

	if err := c.Repository.CreateProduct(r.Context(), command); err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteCreated(w, command.Product, nil)
}

func parseCreateProductRequest(r *http.Request) (*CreateProductCommand, *goutils.ValidationError) {
	validationErrors := map[string]string{}

	product := &Product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &CreateProductCommand{product}, nil
}
