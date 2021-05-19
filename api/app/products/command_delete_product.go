package products

import (
	"net/http"
	"products-api/app/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// DeleteProduct command
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseDeleteProductRequest(r)

	if err != nil {
		c.logError(err)
		utils.WriteValidationError(w, err)
		return
	}

	if err := c.repository.deleteProduct(r.Context(), command); err != nil {
		c.logError(err)
		switch err {
		case utils.ErrNotFound:
			utils.WriteNotFound(w, nil)
		default:
			utils.WriteInternalError(w, nil)
		}
		return
	}

	utils.WriteNoContent(w)
}

func parseDeleteProductRequest(r *http.Request) (*deleteProductCommand, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[utils.FieldID])
	if err != nil {
		errorMap[utils.FieldID] = utils.ErrCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, utils.NewValidationError(errorMap)
	}

	return &deleteProductCommand{id}, nil
}
