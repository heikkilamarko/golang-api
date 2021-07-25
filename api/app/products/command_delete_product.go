package products

import (
	"net/http"
	"products-api/app/utils"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
)

// DeleteProduct command
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseDeleteProductRequest(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.repository.deleteProduct(r.Context(), command); err != nil {
		c.logError(err)
		switch err {
		case utils.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteNoContent(w)
}

func parseDeleteProductRequest(r *http.Request) (*deleteProductCommand, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[utils.FieldID])
	if err != nil {
		errorMap[utils.FieldID] = utils.ErrCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &deleteProductCommand{id}, nil
}
