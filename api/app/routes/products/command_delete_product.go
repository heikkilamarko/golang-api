package products

import (
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// DeleteProduct command
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseDeleteProductRequest(r)

	if err != nil {
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.Repository.DeleteProduct(r.Context(), command); err != nil {
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

func parseDeleteProductRequest(r *http.Request) (*DeleteProductCommand, error) {
	validationErrors := map[string]string{}

	id, err := goutils.GetRequestVarInt(r, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &DeleteProductCommand{id}, nil
}
