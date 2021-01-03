package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"

	"github.com/heikkilamarko/goutils"
)

// CreateProduct command
func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateProductRequest(r)

	if err != nil {
		utils.HandleParseRequestError(err, w)
		return
	}

	if err := c.Repository.CreateProduct(r.Context(), command); err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteCreated(w, command.Product, nil)
}

func parseCreateProductRequest(r *http.Request) (*CreateProductCommand, error) {
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
