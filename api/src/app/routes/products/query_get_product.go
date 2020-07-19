package products

import (
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// GetProduct query
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	query, verr := parseGetProductRequest(r)

	if verr != nil {
		goutils.WriteBadRequest(w, verr.ValidationErrors)
		return
	}

	product, err := c.Repository.GetProduct(r.Context(), query)

	if err != nil {
		switch err {
		case goutils.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, product, nil)
}

func parseGetProductRequest(r *http.Request) (*GetProductQuery, *goutils.ValidationError) {
	validationErrors := map[string]string{}

	id, err := goutils.GetRequestVarInt(r, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &GetProductQuery{id}, nil
}
