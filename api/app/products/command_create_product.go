package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/utils"

	"github.com/heikkilamarko/goutils"
)

// CreateProduct command
func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateProductRequest(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.repository.createProduct(r.Context(), command); err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteCreated(w, command.Product, nil)
}

func parseCreateProductRequest(r *http.Request) (*createProductCommand, error) {
	errorMap := map[string]string{}

	product := &product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		errorMap[utils.FieldRequestBody] = utils.ErrCodeInvalidRequestBody
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &createProductCommand{product}, nil
}
