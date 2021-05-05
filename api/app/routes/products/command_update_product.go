package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
)

// UpdateProduct command
func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseUpdateProductRequest(r)

	if err != nil {
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.repository.updateProduct(r.Context(), command); err != nil {
		switch err {
		case goutils.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, command.Product, nil)
}

func parseUpdateProductRequest(r *http.Request) (*updateProductCommand, error) {
	validationErrors := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[constants.FieldID])
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	product := &product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	if id != product.ID {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &updateProductCommand{product}, nil
}
